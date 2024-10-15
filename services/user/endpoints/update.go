package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/modules/encrypt"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/Capture-411/services-opportunity/services/role"
	"gorm.io/gorm"
)

/*
 * @apiDefine: UpdateUserRequest
 */
type UpdateUserRequest struct {
	Name                 string `json:"name,omitempty" openapi:"example:John;type:string;"`
	Title                string `json:"title,omitempty" openapi:"example:Mr.;type:string;"`
	Phone                string `json:"phone,omitempty" openapi:"example:+1-154-8754-25;type:string;"`
	IDCode               string `json:"id_code,omitempty" openapi:"example:123456789;type:string;"`
	Password             string `json:"password,omitempty" openapi:"example:5f4dcc3b5aa765d61d8327deb882cf99;type:string;"`
	LastName             string `json:"last_name,omitempty" openapi:"example:Doe;type:string;"`
	Username             string `json:"username,omitempty" openapi:"example:johndoe;type:string;"`
	DateOfBirth          string `json:"date_of_birth,omitempty" openapi:"example:1990-01-01;type:date"`
	Gender               string `json:"gender,omitempty" openapi:"example:male;type:string"`
	Address              string `json:"address,omitempty" openapi:"example:123 Main St.;type:string"`
	ZipCode              string `json:"zip_code,omitempty" openapi:"example:12345;type:string"`
	State                string `json:"state,omitempty" openapi:"example:NY;type:string"`
	TownCity             string `json:"town_city,omitempty" openapi:"example:New York;type:string"`
	Country              string `json:"country,omitempty" openapi:"example:USA;type:string"`
	CompanyName          string `json:"company_name,omitempty" openapi:"example:ABC;type:string"`
	Roles                []int  `json:"roles,omitempty" openapi:"$ref:Role;example:[1,2,3];type:array"`
	ContributorID        int64  `json:"contributor_id,omitempty" openapi:"example:1;type:integer"`
	SuspendedAt          bool   `json:"suspended_at,omitempty" openapi:"example:true;type:boolean"`
	EmailVerifiedAt      bool   `json:"email_verified_at,omitempty" openapi:"example:true;type:boolean"`
	PhoneVerifiedAt      bool   `json:"phone_verified_at,omitempty" openapi:"example:true;type:boolean"`
	ReferralCode         string `json:"referral_code,omitempty" openapi:"example:123456789;type:string"`
	ShouldChangePassword bool   `json:"should_change_password,omitempty" openapi:"example:true;type:boolean"`
}

func (r *UpdateUserRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"phone":                  gocriteria.New("phone").Optional(),
		"title":                  gocriteria.New("title").Optional(),
		"name":                   gocriteria.New("name").Required().MinLength(2).MaxLength(200),
		"id_code":                gocriteria.New("id_code").Optional(),
		"last_name":              gocriteria.New("last_name").Required().MinLength(2).MaxLength(100),
		"username":               gocriteria.New("username").Optional(),
		"password":               gocriteria.New("password").Optional(),
		"roles":                  gocriteria.New("roles").Optional(),
		"gender":                 gocriteria.New("gender").Optional(),
		"address":                gocriteria.New("address").Optional(),
		"zip_code":               gocriteria.New("zip_code").Optional(),
		"state":                  gocriteria.New("state").Optional(),
		"town_city":              gocriteria.New("town_city").Optional(),
		"country":                gocriteria.New("country").Optional(),
		"company_name":           gocriteria.New("company_name").Optional(),
		"contributor_id":         gocriteria.New("contributor_id").Optional(),
		"date_of_birth":          gocriteria.New("date_of_birth").Optional(),
		"suspended_at":           gocriteria.New("suspended_at").Optional(),
		"email_verified_at":      gocriteria.New("email_verified_at").Optional(),
		"phone_verified_at":      gocriteria.New("phone_verified_at").Optional(),
		"referral_code":          gocriteria.New("referral_code").Optional(),
		"should_change_password": gocriteria.New("should_change_password").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"name":                   "Name",
			"title":                  "Title",
			"roles":                  "Roles",
			"phone":                  "Phone",
			"id_code":                "ID Code",
			"last_name":              "Lastname",
			"username":               "Username",
			"password":               "Password",
			"gender":                 "Gender",
			"address":                "Address",
			"zip_code":               "Zip Code",
			"state":                  "State",
			"town_city":              "Town/City",
			"country":                "Country",
			"company_name":           "Company Name",
			"contributor_id":         "Contributor ID",
			"date_of_birth":          "Birth Date",
			"suspended_at":           "Suspended at",
			"email_verified_at":      "Email Verified at",
			"phone_verified_at":      "Phone Verified at",
			"referral_code":          "Referral Code",
			"should_change_password": "Should Change Password",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

/*
 * @apiDefine: UpdateUserRequestParam
 */
type UpdateUserRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateUserRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) isEnoughUpdateData(input UpdateUserRequest) bool {
	if input.Name != "" && input.LastName != "" && input.Phone != "" &&
		(input.DateOfBirth != "") {
		return true
	}
	return false
}

func (s *service) Update(ctx context.Context, id string, input UpdateUserRequest) (models.User, response.ErrorResponse) {
	var user models.User
	tx := s.db.WithContext(ctx).Begin() // Begin a new transaction
	if tx.Error != nil {
		s.logger.With(ctx).Error(tx.Error)
		return models.User{}, response.GormErrorResponse(tx.Error, "An error occurred on the server")
	}

	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			s.logger.With(ctx).Error(tx.Error)
			tx.Rollback() // Rollback the transaction if there's any error or panic
		} else {
			err := tx.Commit().Error // Commit the transaction if there are no errors
			if err != nil {
				s.logger.With(ctx).Error(err)
			}
		}
	}()

	if !policy.CanUpdateUser(ctx) {
		s.logger.With(ctx).Error("You do not have permission to access this user")
		return models.User{}, response.ErrorForbidden("You do not have permission to access this user")
	}

	err := tx.Preload("Roles").First(&user, "id = ?", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.User{}, response.GormErrorResponse(err, "Error in finding the user")
	}

	profileCompletedAt := dtp.NullTime{Valid: s.isEnoughUpdateData(input), Time: time.Now()}

	if input.Password != "" {
		var hashedPassword string
		hashedPassword, err = encrypt.HashPassword(input.Password)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return models.User{}, response.ErrorInternalServerError("An error occurred on the server")
		}
		user.Password = hashedPassword
	}

	// Check unique fields to avoid duplication
	if input.Username != "" {
		var tmpUser models.User
		err = tx.First(&tmpUser, "username = ? AND id != ?", input.Username, id).Error
		if err != gorm.ErrRecordNotFound && err != nil {
			s.logger.With(ctx).Error(err)
			return models.User{}, response.GormErrorResponse(err, "Error in finding the user")
		}
		if err == nil && tmpUser.ID != user.ID {
			s.logger.With(ctx).Error("The entered username is duplicated")
			return models.User{}, response.GormErrorResponse(nil, "The entered username is duplicated")
		}
	}

	if input.Phone != "" {
		var tmpUser models.User
		err = tx.First(&tmpUser, "phone = ? AND id != ?", input.Phone, id).Error
		if err != gorm.ErrRecordNotFound && err != nil {
			s.logger.With(ctx).Error(err)
			return models.User{}, response.GormErrorResponse(err, "Error in finding the user")
		}
		if err == nil && tmpUser.ID != user.ID {
			s.logger.With(ctx).Error("The entered phone number is duplicated")
			return models.User{}, response.GormErrorResponse(nil, "The entered phone number is duplicated")
		}
	}

	if input.IDCode != "" {
		var tmpUser models.User
		err = tx.First(&tmpUser, "id_code = ? AND id != ?", input.IDCode, id).Error
		if err != gorm.ErrRecordNotFound && err != nil {
			s.logger.With(ctx).Error(err)
			return models.User{}, response.GormErrorResponse(err, "Error in finding the user")
		}
		if err == nil && tmpUser.ID != user.ID {
			s.logger.With(ctx).Error("The entered national ID code is duplicated")
			return models.User{}, response.GormErrorResponse(nil, "The entered national ID code is duplicated")
		}
	}

	var dateOfBirth time.Time
	if input.DateOfBirth != "" {
		formats := []string{
			"2006-01-02T15:04:05.000Z",
			"Mon Jan 02 2006 15:04:05 MST-0700 (Standard Time)",
			"Mon Jan 02 2006 15:04:05 GMT-0700 (Standard Time)",
			"2006-01-02",
			"2006-01-02T15:04:05Z07:00",
		}
		var parseErr error
		for _, format := range formats {
			dateOfBirth, parseErr = time.Parse(format, input.DateOfBirth)
			if parseErr == nil {
				break
			}
		}
		if parseErr != nil {
			s.logger.With(ctx).Error(parseErr)
			return models.User{}, response.ErrorBadRequest(nil, "Invalid date of birth")
		}
	}

	oldReferralCode := user.ReferralCode
	user.ReferralCode = input.ReferralCode

	user.Name = input.Name
	user.Title = input.Title
	user.Phone = dtp.NullString{
		String: input.Phone,
		Valid:  input.Phone != "",
	}

	user.IDCode = dtp.NullString{
		String: input.IDCode,
		Valid:  input.IDCode != "",
	}
	user.LastName = input.LastName
	user.Username = dtp.NullString{
		String: input.Username,
		Valid:  input.Username != "",
	}
	user.Gender = input.Gender
	user.Address = input.Address
	user.ZipCode = input.ZipCode
	user.State = input.State
	user.TownCity = input.TownCity
	user.Country = input.Country
	user.CompanyName = input.CompanyName
	user.DateOfBirth = dtp.NullTime{
		Time:  dateOfBirth,
		Valid: !dateOfBirth.IsZero(),
	}
	user.ContributorID = dtp.NullInt64{
		Int64: input.ContributorID,
		Valid: input.ContributorID != 0,
	}
	user.SuspendedAt = dtp.NullTime{
		Valid: input.SuspendedAt,
		Time:  time.Now(),
	}
	user.EmailVerifiedAt = dtp.NullTime{
		Valid: input.EmailVerifiedAt,
		Time:  time.Now(),
	}
	user.PhoneVerifiedAt = dtp.NullTime{
		Valid: input.PhoneVerifiedAt,
		Time:  time.Now(),
	}
	user.ProfileCompletedAt = profileCompletedAt

	user.ShouldChangePassword = input.ShouldChangePassword

	// Update roles if provided
	if len(input.Roles) != 0 {
		var roles []*models.Role
		err = tx.Find(&roles, input.Roles).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			return user, response.GormErrorResponse(err, "Error in finding roles")
		}

		// Check if the user is an admin and wants to remove the admin role
		Id := policy.ExtractIdClaim(ctx)
		userID := fmt.Sprintf("%v", user.ID)
		var adminRole *models.Role
		for _, r := range roles {
			if r.Title == role.Admin {
				adminRole = r
				break
			}
		}

		for _, r := range user.Roles {
			if r.Title == role.Admin && adminRole == nil && userID == Id {
				err := response.ErrorForbidden("You do not have permission to remove the admin role")
				s.logger.With(ctx).Error(err)
				return models.User{}, err
			}
		}

		err = tx.Model(&user).Association("Roles").Replace(roles)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return models.User{}, response.GormErrorResponse(err, "Error in saving the user")
		}
	}

	err = tx.Save(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.User{}, response.GormErrorResponse(err, "Error in saving the user")
	}

	// Update referred with code if referral code changes
	if oldReferralCode != input.ReferralCode {
		err = tx.Model(&models.User{}).Where("referred_with_code = ?", oldReferralCode).Update("referred_with_code", input.ReferralCode).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			tx.Rollback()
			return models.User{}, response.GormErrorResponse(err, "Error in updating referred with code")
		}
	}

	return user, response.ErrorResponse{}
}
