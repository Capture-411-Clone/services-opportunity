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
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateUserAccountRequest
 */
type UpdateUserAccountRequest struct {
	Name        string `json:"name,omitempty" openapi:"example:my name;type:string"`
	Title       string `json:"title,omitempty" openapi:"example:my title;type:string"`
	Phone       string `json:"phone,omitempty" openapi:"example:+1-878-24-4587;type:string"`
	IDCode      string `json:"id_code,omitempty" openapi:"example:ABCD12345678;type:string"`
	LastName    string `json:"last_name,omitempty" openapi:"example:my last name;type:string"`
	Username    string `json:"username,omitempty" openapi:"example:7209311196;type:string"`
	DateOfBirth string `json:"date_of_birth,omitempty" openapi:"example:2020-01-01;type:datetime"`
	Gender      string `json:"gender,omitempty" openapi:"example:MALE;type:string"`
	Address     string `json:"address,omitempty" openapi:"example:my address;type:string"`
	ZipCode     string `json:"zip_code,omitempty" openapi:"example:12345;type:string"`
	State       string `json:"state,omitempty" openapi:"example:NY;type:string"`
	TownCity    string `json:"town_city,omitempty" openapi:"example:New York;type:string"`
	Country     string `json:"country,omitempty" openapi:"example:USA;type:string"`
	CompanyName string `json:"company_name,omitempty" openapi:"example:my company name;type:string"`
}

func (s *service) isEnoughUpdateAccountData(user models.User, input UpdateUserAccountRequest) bool {
	if input.Name != "" && input.LastName != "" && input.Username != "" && input.Phone != "" && input.IDCode != "" &&
		input.Gender != "" && (input.DateOfBirth != "") && user.PhoneVerifiedAt.Valid {
		return true
	}
	return false
}

func (req *UpdateUserAccountRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"name":          gocriteria.New("name").Required().MinMaxLength(2, 200),
		"Title":         gocriteria.New("title").Optional(),
		"phone":         gocriteria.New("phone").Optional(),
		"id_code":       gocriteria.New("id_code").Optional(),
		"username":      gocriteria.New("username").Required().MinMaxLength(2, 200),
		"last_name":     gocriteria.New("last_name").Required().MinMaxLength(2, 200),
		"gender":        gocriteria.New("gender").Optional(),
		"date_of_birth": gocriteria.New("date_of_birth").Optional(),
		"address":       gocriteria.New("address").Optional(),
		"zip_code":      gocriteria.New("zip_code").Optional(),
		"state":         gocriteria.New("state").Optional(),
		"town_city":     gocriteria.New("town_city").Optional(),
		"country":       gocriteria.New("country").Optional(),
		"company_name":  gocriteria.New("company_name").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"name":          "Name",
			"title":         "Title",
			"id_code":       "National ID",
			"phone":         "Phone",
			"username":      "Username",
			"last_name":     "Lastname",
			"gender":        "Gender",
			"date_of_birth": "Birth Date",
			"address":       "Address",
			"zip_code":      "Zip Code",
			"state":         "State",
			"town_city":     "Town/City",
			"country":       "Country",
			"company_name":  "Company Name",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, req)
	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) UpdateAccount(ctx context.Context, input UpdateUserAccountRequest) (models.User, response.ErrorResponse) {
	var user models.User
	Id := policy.ExtractIdClaim(ctx)
	err := s.db.WithContext(ctx).First(&user, "id", Id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return user, response.GormErrorResponse(err, "Error in finding the user")
	}

	if !policy.CanUpdateAccount(ctx, user) {
		return user, response.ErrorForbidden("You do not have permission to access this user")
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

	phoneVerifiedAt := user.PhoneVerifiedAt
	if input.Phone != user.Phone.String {
		phoneVerifiedAt = dtp.NullTime{Valid: false, Time: time.Now()}
	}

	profileCompletedAt := dtp.NullTime{Valid: s.isEnoughUpdateAccountData(user, input), Time: time.Now()}

	user.Name = input.Name
	user.Title = input.Title
	user.IDCode = dtp.NullString{String: input.IDCode, Valid: input.IDCode != ""}
	user.Phone = dtp.NullString{
		String: input.Phone,
		Valid:  input.Phone != "",
	}
	user.PhoneVerifiedAt = phoneVerifiedAt
	user.Username = dtp.NullString{String: input.Username, Valid: input.Username != ""}
	user.LastName = input.LastName
	user.Gender = input.Gender
	user.DateOfBirth = dtp.NullTime{
		Time:  dateOfBirth,
		Valid: !dateOfBirth.IsZero(),
	}
	user.Address = input.Address
	user.ZipCode = input.ZipCode
	user.State = input.State
	user.TownCity = input.TownCity
	user.Country = input.Country
	user.CompanyName = input.CompanyName
	user.ProfileCompletedAt = profileCompletedAt

	err = s.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return user, response.GormErrorResponse(err, "Error in saving the user")
	}

	return user, response.ErrorResponse{}
}
