package endpoints

import (
	"context"
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
)

/*
 * @apiDefine: CreateUserRequest
 */
type CreateUserRequest struct {
	Name                 string `json:"name,omitempty" openapi:"example:John;type:string;"`
	Title                string `json:"title,omitempty" openapi:"example:Mr.;type:string;"`
	Email                string `json:"email,omitempty" openapi:"example:Vr5kH@example.com;type:string;"`
	Phone                string `json:"phone,omitempty" openapi:"example:+1-154-8754-25;type:string;"`
	IDCode               string `json:"id_code,omitempty" openapi:"example:123456789;type:string;"`
	Password             string `json:"password,omitempty" openapi:"example:password;type:string;"`
	LastName             string `json:"last_name,omitempty" openapi:"example:Doe;type:string;"`
	Username             string `json:"username,omitempty" openapi:"example:johndoe;type:string;"`
	DateOfBirth          string `json:"date_of_birth,omitempty" openapi:"example:1990-01-01;type:date;"`
	Gender               string `json:"gender,omitempty" openapi:"example:male;type:string;"`
	Address              string `json:"address,omitempty" openapi:"example:123 Main St.;type:string;"`
	ZipCode              string `json:"zip_code,omitempty" openapi:"example:12345;type:string;"`
	State                string `json:"state,omitempty" openapi:"example:NY;type:string;"`
	TownCity             string `json:"town_city,omitempty" openapi:"example:New York;type:string;"`
	Country              string `json:"country,omitempty" openapi:"example:USA;type:string;"`
	CompanyName          string `json:"company_name,omitempty" openapi:"example:ABC;type:string;"`
	Roles                []int  `json:"roles,omitempty" openapi:"example:[1,2];type:array;"`
	SuspendedAt          bool   `json:"suspended_at,omitempty" openapi:"example:true;type:boolean;"`
	ContributorID        int64  `json:"contributor_id,omitempty" openapi:"example:1;type:integer;"`
	EmailVerifiedAt      bool   `json:"email_verified_at,omitempty" openapi:"example:true;type:boolean;"`
	ReferralCode         string `json:"referral_code,omitempty" openapi:"example:123456789;type:string;"`
	PhoneVerifiedAt      bool   `json:"phone_verified_at,omitempty" openapi:"example:true;type:boolean;"`
	ShouldChangePassword bool   `json:"should_change_password,omitempty" openapi:"example:true;type:boolean;"`
}

func (r *CreateUserRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		// TODO: username should include not only number to prevent reserved phone hack
		"username":               gocriteria.New("username").Optional(),
		"phone":                  gocriteria.New("phone").Optional(),
		"title":                  gocriteria.New("title").Optional(),
		"name":                   gocriteria.New("name").Required().MinLength(2).MaxLength(200),
		"email":                  gocriteria.New("email").Email().Required(),
		"id_code":                gocriteria.New("id_code").Optional(),
		"last_name":              gocriteria.New("last_name").Required().MinLength(2).MaxLength(100),
		"password":               gocriteria.New("password").MinLength(8).Required(),
		"roles":                  gocriteria.New("roles").Optional(),
		"gender":                 gocriteria.New("gender").Optional(),
		"address":                gocriteria.New("address").Optional(),
		"zip_code":               gocriteria.New("zip_code").Optional(),
		"state":                  gocriteria.New("state").Optional(),
		"town_city":              gocriteria.New("town_city").Optional(),
		"country":                gocriteria.New("country").Optional(),
		"company_name":           gocriteria.New("company_name").Optional(),
		"date_of_birth":          gocriteria.New("date_of_birth").Optional(),
		"suspended_at":           gocriteria.New("suspended_at").Optional(),
		"contributor_id":         gocriteria.New("contributor_id").Optional(),
		"email_verified_at":      gocriteria.New("email_verified_at").Optional(),
		"referral_code":          gocriteria.New("referral_code").Optional(),
		"phone_verified_at":      gocriteria.New("phone_verified_at").Optional(),
		"should_change_password": gocriteria.New("should_change_password").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"name":                   "Name",
			"title":                  "Title",
			"phone":                  "phone",
			"email":                  "Email",
			"roles":                  "Roles",
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
			"date_of_birth":          "Birth Date",
			"suspended_at":           "Suspended at",
			"contributor_id":         "Contributor ID",
			"email_verified_at":      "Email Verified at",
			"referral_code":          "Referral Code",
			"phone_verified_at":      "Phone Verified at",
			"should_change_password": "Should Change Password",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) isEnoughCreateData(input CreateUserRequest) bool {
	if input.Name != "" && input.LastName != "" && input.Phone != "" && input.Email != "" &&
		(input.DateOfBirth != "") && input.PhoneVerifiedAt && input.EmailVerifiedAt {
		return true
	}
	return false
}

func (s *service) Create(ctx context.Context, input CreateUserRequest) (models.User, response.ErrorResponse) {
	if !policy.CanCreateUser(ctx) {
		s.logger.With(ctx).Error("You do not have permission to access this user")
		return models.User{}, response.ErrorForbidden(nil, "You do not have permission to access this user")
	}

	profileCompletedAt := dtp.NullTime{Valid: s.isEnoughCreateData(input), Time: time.Now()}

	var hashedPassword string
	hashedPassword, err := encrypt.HashPassword(input.Password)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.User{}, response.ErrorInternalServerError(nil, "An error occurred on the server")
	}

	var rls []*models.Role
	if len(input.Roles) != 0 {
		err := s.db.Find(&rls, input.Roles).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			return models.User{}, response.ErrorBadRequest(nil, "Invalid roles sent")
		}
	} else {
		var r models.Role
		err := s.db.First(&r, "title = ?", role.User).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			return models.User{}, response.GormErrorResponse(nil, "Faild Setting Role to User")
		}
		rls = append(rls, &r)
	}

	var dateOfBirth time.Time
	if input.DateOfBirth != "" {
		formats := []string{
			"2006-01-02T15:04:05.000Z",
			"Mon Jan 02 2006 15:04:05 MST-0700 (Standard Time)",
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

	user := models.User{
		ReferralCode: input.ReferralCode,

		Name:  input.Name,
		Title: input.Title,
		Phone: dtp.NullString{
			String: input.Phone,
			Valid:  input.Phone != "",
		},
		ContributorID: dtp.NullInt64{
			Int64: input.ContributorID,
			Valid: input.ContributorID != 0,
		},
		LastName:    input.LastName,
		Password:    hashedPassword,
		Roles:       rls,
		Gender:      input.Gender,
		Address:     input.Address,
		ZipCode:     input.ZipCode,
		State:       input.State,
		TownCity:    input.TownCity,
		Country:     input.Country,
		CompanyName: input.CompanyName,
		DateOfBirth: dtp.NullTime{
			Time:  dateOfBirth,
			Valid: dateOfBirth != time.Time{},
		},
		Username: dtp.NullString{
			String: input.Username,
			Valid:  input.Username != "",
		},
		Email: dtp.NullString{
			String: input.Email,
			Valid:  input.Email != "",
		},
		IDCode: dtp.NullString{
			String: input.IDCode,
			Valid:  input.IDCode != "",
		},
		SuspendedAt: dtp.NullTime{
			Valid: input.SuspendedAt,
			Time:  time.Now(),
		},
		EmailVerifiedAt: dtp.NullTime{
			Valid: input.EmailVerifiedAt,
			Time:  time.Now(),
		},
		PhoneVerifiedAt: dtp.NullTime{
			Valid: input.PhoneVerifiedAt,
			Time:  time.Now(),
		},
		ProfileCompletedAt: profileCompletedAt,

		ShouldChangePassword: input.ShouldChangePassword,
	}
	err = s.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.User{}, response.GormErrorResponse(err, "An error occurred while registering the user")
	}
	return user, response.ErrorResponse{}
}
