package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/exp"
	"github.com/Capture-411/services-opportunity/kit/filter"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UserQueryFilterType
 */
type UserQueryFilterType struct {
	RoleID             filter.FilterValue[string] `json:"role_id,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Name               filter.FilterValue[string] `json:"name,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Phone              filter.FilterValue[string] `json:"phone,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Email              filter.FilterValue[string] `json:"email,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	IDCode             filter.FilterValue[string] `json:"id_code,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	LastName           filter.FilterValue[string] `json:"last_name,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Username           filter.FilterValue[string] `json:"username,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Password           filter.FilterValue[string] `json:"password,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	RoleIDs            filter.FilterValue[string] `json:"role_ids,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Gender             filter.FilterValue[string] `json:"gender,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Country            filter.FilterValue[string] `json:"country,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	CompanyName        filter.FilterValue[string] `json:"company_name,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	State              filter.FilterValue[string] `json:"state,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	TownCity           filter.FilterValue[string] `json:"town_city,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	DateOfBirth        filter.FilterValue[string] `json:"date_of_birth,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	SuspendedAt        filter.FilterValue[string] `json:"suspended_at,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	EmailVerifiedAt    filter.FilterValue[string] `json:"email_verified_at,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	PhoneVerifiedAt    filter.FilterValue[string] `json:"phone_verified_at,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	ProfileCompletedAt filter.FilterValue[string] `json:"profile_completed_at,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	ReferralCode       filter.FilterValue[string] `json:"referral_code,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	ReferredWithCode   filter.FilterValue[string] `json:"referred_with_code,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	CreatedAt          filter.FilterValue[string] `json:"created_at" openapi:"nullable;$ref:FilterValueString;"`
}

/*
 * @apiDefine: UserQueryRequestParams
 */
type UserQueryRequestParams struct {
	ID      string              `json:"id,omitempty" openapi:"example:1;required;pattern:^[0-9]+$;in:query"`
	Page    string              `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage string              `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order   string              `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy string              `json:"order_by,omitempty" openapi:"example:type;in:query"`
	Query   string              `json:"query,omitempty" openapi:"pattern:^[0-9]+$;in:query"`
	Filters UserQueryFilterType `json:"filters,omitempty" openapi:"$ref:UserQueryFilterType;type:object"`
}

func (data *UserQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":       gocriteria.New("id").Optional(),
		"page":     gocriteria.New("page").Optional(),
		"per_page": gocriteria.New("per_page").Optional(),
		"query":    gocriteria.New("query").Optional(),
		"order":    gocriteria.New("order").Optional(),
		"order_by": gocriteria.New("order_by").Optional(),
		"filters": gocriteria.Schema{
			"name": gocriteria.Schema{
				"op":    gocriteria.New("filter.name.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.name.value").Optional(),
			},
			"phone": gocriteria.Schema{
				"op":    gocriteria.New("filter.phone.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.phone.value").Optional(),
			},
			"email": gocriteria.Schema{
				"op":    gocriteria.New("filter.email.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.email.value").Optional(),
			},
			"id_code": gocriteria.Schema{
				"op":    gocriteria.New("filter.id_code.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.id_code.value").Optional(),
			},
			"last_name": gocriteria.Schema{
				"op":    gocriteria.New("filter.last_name.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.last_name.value").Optional(),
			},
			"username": gocriteria.Schema{
				"op":    gocriteria.New("filter.username.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.username.value").Optional(),
			},
			"role_ids": gocriteria.Schema{
				"op":    gocriteria.New("filter.role_ids.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.role_ids.value").Optional(),
			},
			"gender": gocriteria.Schema{
				"op":    gocriteria.New("filter.gender.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.gender.value").Optional(),
			},
			"country": gocriteria.Schema{
				"op":    gocriteria.New("filter.country.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.country.value").Optional(),
			},
			"company_name": gocriteria.Schema{
				"op":    gocriteria.New("filter.company_name.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.company_name.value").Optional(),
			},
			"state": gocriteria.Schema{
				"op":    gocriteria.New("filter.state.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.state.value").Optional(),
			},
			"town_city": gocriteria.Schema{
				"op":    gocriteria.New("filter.town_city.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.town_city.value").Optional(),
			},
			"date_of_birth": gocriteria.Schema{
				"op":    gocriteria.New("filter.date_of_birth.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.date_of_birth.value").Optional(),
			},
			"suspended_at": gocriteria.Schema{
				"op":    gocriteria.New("filter.suspended_at.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.suspended_at.value").Optional(),
			},
			"email_verified_at": gocriteria.Schema{
				"op":    gocriteria.New("filter.email_verified_at.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.email_verified_at.value").Optional(),
			},
			"phone_verified_at": gocriteria.Schema{
				"op":    gocriteria.New("filter.phone_verified_at.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.phone_verified_at.value").Optional(),
			},
			"profile_completed_at": gocriteria.Schema{
				"op":    gocriteria.New("filter.profile_completed_at.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.profile_completed_at.value").Optional(),
			},
			"referral_code": gocriteria.Schema{
				"op":    gocriteria.New("filter.referral_code.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.referral_code.value").Optional(),
			},
			"referred_with_code": gocriteria.Schema{
				"op":    gocriteria.New("filter.referred_with_code.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.referred_with_code.value").Optional(),
			},
			"created_at": gocriteria.Schema{
				"op":    gocriteria.New("filter.created_at.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.created_at.value").Optional(),
			},
		},
	}

	errs := gocriteria.ValidateQueries(r, schema, data)
	if len(errs) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errs)
		return dumpedErrors
	}

	return nil
}

func makeFilters(params UserQueryRequestParams) []string {
	var where []string

	if params.ID != "" {
		where = append(where, fmt.Sprintf("id = %s", params.ID))
	}

	// if params.Filters.RoleID.Op != "" {
	// 	opValue := filter.GetDBOperatorAndValue(params.Filters.RoleID.Op, params.Filters.RoleID.Value)
	// 	val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
	// 	where = append(where, fmt.Sprintf("role_id %s %s", opValue.Operator, val))
	// }
	if params.Filters.Name.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Name.Op, params.Filters.Name.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("name %s %s", opValue.Operator, val))
	}
	if params.Filters.Phone.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Phone.Op, params.Filters.Phone.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("phone %s %s", opValue.Operator, val))
	}
	if params.Filters.Email.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Email.Op, params.Filters.Email.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("email %s %s", opValue.Operator, val))
	}
	if params.Filters.IDCode.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.IDCode.Op, params.Filters.IDCode.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("id_code %s %s", opValue.Operator, val))
	}
	if params.Filters.LastName.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.LastName.Op, params.Filters.LastName.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("last_name %s %s", opValue.Operator, val))
	}
	if params.Filters.Username.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Username.Op, params.Filters.Username.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("username %s %s", opValue.Operator, val))
	}
	if params.Filters.Password.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Password.Op, params.Filters.Password.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("password %s %s", opValue.Operator, val))
	}
	if params.Filters.RoleIDs.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.RoleIDs.Op, params.Filters.RoleIDs.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("role_ids %s %s", opValue.Operator, val))
	}
	if params.Filters.Gender.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Gender.Op, params.Filters.Gender.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("gender %s %s", opValue.Operator, val))
	}
	if params.Filters.Country.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Country.Op, params.Filters.Country.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("country %s %s", opValue.Operator, val))
	}
	if params.Filters.CompanyName.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CompanyName.Op, params.Filters.CompanyName.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("company_name %s %s", opValue.Operator, val))
	}
	if params.Filters.State.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.State.Op, params.Filters.State.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("state %s %s", opValue.Operator, val))
	}
	if params.Filters.TownCity.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.TownCity.Op, params.Filters.TownCity.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("town_city %s %s", opValue.Operator, val))
	}
	if params.Filters.DateOfBirth.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.DateOfBirth.Op, params.Filters.DateOfBirth.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("date_of_birth %s %s", opValue.Operator, val))
	}
	if params.Filters.SuspendedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.SuspendedAt.Op, params.Filters.SuspendedAt.Value)
		// transform boolean to to null or not null date time
		if opValue.Value != "" {
			transformedValue := exp.TerIf(opValue.Value == "true", "IS NOT NULL", "IS NULL")
			where = append(where, fmt.Sprintf("suspended_at %s", transformedValue))
		}
	}
	if params.Filters.EmailVerifiedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.EmailVerifiedAt.Op, params.Filters.EmailVerifiedAt.Value)
		// transform boolean to to null or not null date time
		if opValue.Value != "" {
			transformedValue := exp.TerIf(opValue.Value == "true", "IS NOT NULL", "IS NULL")
			where = append(where, fmt.Sprintf("email_verified_at %s", transformedValue))
		}
	}
	if params.Filters.PhoneVerifiedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.PhoneVerifiedAt.Op, params.Filters.PhoneVerifiedAt.Value)
		// transform boolean to to null or not null date time
		if opValue.Value != "" {
			transformedValue := exp.TerIf(opValue.Value == "true", "IS NOT NULL", "IS NULL")
			where = append(where, fmt.Sprintf("phone_verified_at %s", transformedValue))
		}
	}
	if params.Filters.ProfileCompletedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.ProfileCompletedAt.Op, params.Filters.ProfileCompletedAt.Value)
		// transform boolean to to null or not null date time
		if opValue.Value != "" {
			transformedValue := exp.TerIf(opValue.Value == "true", "IS NOT NULL", "IS NULL")
			where = append(where, fmt.Sprintf("profile_completed_at %s", transformedValue))
		}
	}
	if params.Filters.ReferralCode.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.ReferralCode.Op, params.Filters.ReferralCode.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("referral_code %s %s", opValue.Operator, val))
	}
	if params.Filters.ReferredWithCode.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.ReferredWithCode.Op, params.Filters.ReferredWithCode.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("referred_with_code %s %s", opValue.Operator, val))
	}
	if params.Filters.CreatedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CreatedAt.Op, params.Filters.CreatedAt.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("created_at %s %s", opValue.Operator, val))
	}

	if len(where) > 0 {
		// wrap any where item with () to avoid conflict with other where items
		for i, w := range where {
			where[i] = fmt.Sprintf("(%s)", w)
		}
	}

	return where

}

func (s *service) Query(
	ctx context.Context, offset, limit int, params UserQueryRequestParams) ([]models.User, response.ErrorResponse) {
	var users []models.User

	if !policy.CanQueryUsers(ctx) {
		s.logger.With(ctx).Error("You do not have permission to access this user")
		return users, response.ErrorForbidden("You do not have permission to access this user")
	}
	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order)).
		Preload("Invite").
		Preload("Contributor").
		Preload("Roles")

	where := makeFilters(params)
	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	err := tx.Find(&users).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return users, response.GormErrorResponse(err, "Error in fetching users")
	}

	return users, response.ErrorResponse{}
}
