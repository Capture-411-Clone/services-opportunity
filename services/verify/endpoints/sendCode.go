package endpoints

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"gorm.io/gorm"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/modules/codegen"
	emailtemplates "github.com/Capture-411/services-opportunity/notification-v2/email/email-templates"
	"github.com/gofrs/uuid"
)

/*
 * @apiDefine: SendCodeRequest
 */
type SendCodeRequest struct {
	Phone string `json:"phone,omitempty" openapi:"example:+1-154-8754-25;type:string"`
	Email string `json:"email,omitempty" openapi:"example:christina@example.com;type:string"`
}

func (r *SendCodeRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"phone": gocriteria.New("phone").Optional(),
		"email": gocriteria.New("email").Email().Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"phone": "Phone",
			"email": "Email",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)
	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func checkError(ctx context.Context, logger log.Logger, err error) bool {
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.With(ctx).Error(err)
			// 	return true
		}
		return false
	}
	return false
}

func (s *service) SendCode(ctx context.Context, input SendCodeRequest) (string, response.ErrorResponse) {
	fmt.Println("SendCode")
	if input.Phone == "" && input.Email == "" {
		s.logger.With(ctx).Error("At least one of the phone or email fields must have a value")
		return "", response.ErrorBadRequest(nil, "At least one of the phone or email fields must have a value")
	}

	if config.AppConfig.Environment != "development" {
		if input.Phone != "" {
			var lastPhoneCode models.Verification
			err := s.db.WithContext(ctx).Last(&lastPhoneCode, "phone", input.Phone).Error
			if checkError(ctx, s.logger, err) {
				s.logger.With(ctx).Error(err)
				return "", response.GormErrorResponse(err, "An error occurred while finding the code")
			}

			if !lastPhoneCode.Expired() {
				return "", response.ErrorResponse{}
				// seconds := int(time.Until(lastPhoneCode.ExpiresAt).Seconds())
				// er := response.GormErrorResponse(err,
				// 	"Sending the code is not possible. Please wait for "+
				// 		fmt.Sprintf("%d", seconds)+
				// 		" seconds ",
				// )
				// return "", er
			}
		}

		if input.Email != "" {
			var lastEmailCode models.Verification
			err := s.db.WithContext(ctx).Last(&lastEmailCode, "email", input.Email).Error
			if checkError(ctx, s.logger, err) {
				s.logger.With(ctx).Error(err)
				return "", response.GormErrorResponse(err, "An error occurred while finding the code")
			}

			if !lastEmailCode.Expired() {
				return "", response.ErrorResponse{}
				// seconds := int(time.Until(lastEmailCode.ExpiresAt).Seconds())
				// er := response.GormErrorResponse(err,
				// 	"Sending the code is not possible. Please wait for "+
				// 		fmt.Sprintf("%d", seconds)+
				// 		" seconds ",
				// )
				// s.logger.With(ctx).Error(err)
				// return "", er
			}
		}

	}

	uid, e := uuid.NewV4()
	if e != nil {
		s.logger.With(ctx).Error(e)
		return "", response.ErrorInternalServerError(e, "An error occurred while creating the code")
	}

	verificationCodeExpireMinutes := config.AppConfig.VerificationCodeExpireMinutes

	verification := models.Verification{
		Code:        codegen.GenerateRandomNumber(),
		Phone:       input.Phone,
		Email:       input.Email,
		ExpiresAt:   time.Now().Add(time.Minute * time.Duration(verificationCodeExpireMinutes)),
		SessionCode: uid.String(),
	}
	err := s.db.WithContext(ctx).Create(&verification).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred while creating the code")
	}
	code := verification.Code

	if config.AppConfig.SendEmail {
		data := emailtemplates.VerificationCodeData{
			Code:             code,
			ResetPasswordURL: "https://capture411.com/auth/reset-password/",
		}
		subject, plainMessage, htmlMessage, err := emailtemplates.VerificationCodeTemplate(data)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorInternalServerError(err, "An error occurred while sending the code")
		}

		if input.Email != "" {
			if config.AppConfig.EmailDriver == "AWS" {
				err = sendCodeWithAWSSES(input.Email, plainMessage)
				if err != nil {
					s.logger.With(ctx).Error(err)
					return "", response.ErrorInternalServerError(err, "sendCodeWithAWSSES, An error occurred while sending the code")
				}
			} else if config.AppConfig.EmailDriver == "CUSTOM" {
				err = s.notifierv2.SendNotification(ctx, []models.User{{Email: dtp.NullString{String: input.Email, Valid: true}}}, nil, []string{"email"}, plainMessage, htmlMessage, subject)
				if err != nil {
					s.logger.With(ctx).Error(err)
					return "", response.ErrorInternalServerError(err, "Sending the code is not possible. Please try again later")
				}
			}
		}
	}

	if config.AppConfig.Environment != "development" {
		code = "successfully sent"
	}
	return code, response.ErrorResponse{}

}

func sendCodeWithAWSSES(Recipient, htmlMsg string) error {
	const (
		Sender  = "no-reply@capture411.com"
		Subject = "Capture411 Verification Code"
		// TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
		CharSet = "UTF-8"
	)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.AppConfig.AwsRegion), // Replace with your AWS region
		Credentials: credentials.NewStaticCredentials(config.AppConfig.AwsAccessKeyId, config.AppConfig.AwsSecretAccessKey, ""),
	})

	if err != nil {
		return err
	}

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlMsg),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(""),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	}

	// Attempt to send the email.
	_, err = svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return err
	}

	fmt.Println("Email Sent to address: " + Recipient)

	return nil
}
