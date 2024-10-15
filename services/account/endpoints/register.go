package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/locales"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/faker"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/modules/encrypt"
	emailtemplates "github.com/Capture-411/services-opportunity/notification-v2/email/email-templates"
	"github.com/Capture-411/services-opportunity/services/role"
)

/*
 * @apiDefine: RegisterRequest
 */
type RegisterRequest struct {
	Bio              string `json:"biography,omitempty" openapi:"example:my bio;type:string"`
	Name             string `json:"name,omitempty" openapi:"example:my name;type:string"`
	Email            string `json:"email,omitempty" openapi:"example:JnY6A@example.com;type:string"`
	Phone            string `json:"phone,omitempty" openapi:"example:+1-878-24-4587;type:string"`
	LastName         string `json:"last_name,omitempty" openapi:"example:my last name;type:string"`
	Password         string `json:"password,omitempty" openapi:"example:1234568;type:string"`
	SessionCode      string `json:"session_code,omitempty" openapi:"example:1234568;type:string"`
	InviteCode       string `json:"invite_code" openapi:"example:invite_code;type:string"`
	ReferredWithCode string `json:"referred_with_code" openapi:"example:referred_with_code;type:string"`
}

func (r *RegisterRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"email":              gocriteria.New("email").Email().Required(),
		"password":           PasswordValidity,
		"name":               gocriteria.New("name").MinLength(2).MaxLength(200).Required(),
		"last_name":          gocriteria.New("last_name").MinLength(2).MaxLength(200).Required(),
		"session_code":       gocriteria.New("session_code").Required(),
		"invite_code":        gocriteria.New("invite_code").AlphaNum(locales.EnUS).Optional(),
		"referred_with_code": gocriteria.New("referred_with_code").AlphaNum(locales.EnUS).Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"name":               "Name",
			"last_name":          "Lastname",
			"email":              "Email",
			"password":           "Password",
			"session_code":       "Session Code",
			"invite_code":        "Invite Code",
			"referred_with_code": "Referred With Code",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Register(ctx context.Context, input RegisterRequest) (LoginResponseData, response.ErrorResponse) {
	var user models.User
	var exists bool
	var res LoginResponseData

	exists, user, err := s.findUser(ctx, input.Email)
	if exists {
		return res, response.ErrorBadRequest(nil, "A user with this username was found")
	}

	if err != nil {
		s.logger.With(ctx).Error(err)
		return res, response.GormErrorResponse(err, "An error occurred")
	}

	if config.AppConfig.RequiredSignUpInviteCode {
		if input.InviteCode == "" {
			return res, response.ErrorBadRequest(nil, "invite code required")
		}

		invite, er := s.useInviteCode(ctx, input.InviteCode)
		if er.StatusCode != 0 {
			s.logger.With(ctx).Error(err)
			return res, er
		}

		registerInvite := models.RegisterInvite{
			HostID: invite.UserID,
			UserID: int(user.ID),
		}

		err := s.db.WithContext(ctx).Create(&registerInvite).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			return res, response.GormErrorResponse(err, "an error occurred between host and invite")
		}
	}

	_, eerr := s.checkAndDeleteVerificationBySessionCodeAndEmail(ctx, input.SessionCode, input.Email)
	if eerr.StatusCode != 0 {
		s.logger.With(ctx).Error(eerr)
		return res, eerr
	}

	var hashedPassword string
	hashedPassword, err = encrypt.HashPassword(input.Password)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return res, response.ErrorInternalServerError(nil, "An error occurred while creating the password")
	}

	var rls []*models.Role
	var r models.Role
	err = s.db.First(&r, "title = ?", role.User).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return res, response.GormErrorResponse(err, "Failed Setting Role to User")
	}
	rls = append(rls, &r)

	user = models.User{
		Name:     input.Name,
		LastName: input.LastName,
		Email: dtp.NullString{
			String: input.Email,
			Valid:  true,
		},
		Roles:            rls,
		Password:         hashedPassword,
		EmailVerifiedAt:  faker.SQLNow(),
		ReferredWithCode: input.ReferredWithCode,
	}
	err = s.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return res, response.GormErrorResponse(err, "An error occurred while registering the user")
	}

	accessToken, err := s.generateAccessTokenJWT(user)
	if err != nil {
		return res, response.ErrorBadRequest(err, "An error occurred while generating the token")
	}

	res = LoginResponseData{
		User:        user,
		AccessToken: accessToken,
	}

	notifData := emailtemplates.UserRegisteredWelcomeData{
		UserFullName: user.Name + " " + user.LastName,
	}
	s.sendRegisterWelcomeNotification(ctx, user, notifData)

	return res, response.ErrorResponse{}
}

func (s *service) sendRegisterWelcomeNotification(ctx context.Context, user models.User, data emailtemplates.UserRegisteredWelcomeData) {
	var targetUsers []models.User
	targetUsers = append(targetUsers, user)

	subject, plainMessage, htmlMessage, err := emailtemplates.UserRegisteredWelcomeTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	subjectShort, messageShort, err := emailtemplates.UserRegisteredWelcomeShortTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	err = s.notifierv2.SendNotification(ctx, targetUsers, nil, []string{"email"}, plainMessage, htmlMessage, subject)
	if err != nil {
		s.logger.Errorf("Error sending email notification: %v\n", err)
	}

	err = s.notifierv2.SendNotification(ctx, targetUsers, nil, []string{"database"}, messageShort, "", subjectShort)
	if err != nil {
		s.logger.Errorf("Error sending database notification: %v\n", err)
	}
}
