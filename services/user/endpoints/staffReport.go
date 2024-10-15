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
 * @apiDefine: StaffReportQueryFilterType
 */
type StaffReportQueryFilterType struct {
	Email filter.FilterValue[string] `json:"email,omitempty" openapi:"nullable;$ref:FilterValueString;"`
}

/*
 * @apiDefine: StaffReportQueryRequestParams
 */
type StaffReportQueryRequestParams struct {
	ID      string                     `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page    string                     `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage string                     `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order   string                     `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy string                     `json:"order_by,omitempty" openapi:"example:title;in:query"`
	Query   string                     `json:"query,omitempty" openapi:"in:query"`
	Filters StaffReportQueryFilterType `json:"filters,omitempty" openapi:"$ref:StaffReportQueryFilterType;in:query"`
}

func (data *StaffReportQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":       gocriteria.New("id").Optional(),
		"page":     gocriteria.New("page").Optional(),
		"per_page": gocriteria.New("per_page").Optional(),
		"query":    gocriteria.New("query").Optional(),
		"order":    gocriteria.New("order").Optional(),
		"order_by": gocriteria.New("order_by").Optional(),
		"filters": gocriteria.Schema{
			"email": gocriteria.Schema{
				"op":    gocriteria.New("filter.email.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.email.value").Optional(),
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

func makeStaffReportFilters(params StaffReportQueryRequestParams) []string {
	var where []string

	if params.ID != "" {
		where = append(where, fmt.Sprintf("users.id = %s", params.ID))
	}

	if params.Filters.Email.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Email.Op, params.Filters.Email.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("users.email %s %s", opValue.Operator, val))
	}

	return where
}

/*
 * @apiDefine: StaffReportData
 */
type StaffReportResponseData struct {
	User              models.User `json:"user"`
	ContributionCount int64       `json:"contribution_count"`
	LastLoginAt       string      `json:"last_login_at"`
}

func (s *service) StaffReport(ctx context.Context, offset, limit int, params StaffReportQueryRequestParams) ([]StaffReportResponseData, response.ErrorResponse) {
	if !policy.CanSeeStaffReport(ctx) {
		s.logger.With(ctx).Error("You do not have permission to see staff report")
		return nil, response.ErrorForbidden("You do not have permission to see the staff report")
	}

	var staffReportData []StaffReportResponseData

	var staff []models.User

	tx := s.db.WithContext(ctx).Where("contributor_id IS NOT NULL").Offset(offset).Limit(limit)
	where := makeStaffReportFilters(params)
	if len(where) > 0 {
		// wrap any where item with () to avoid conflict with other where items
		for i, w := range where {
			where[i] = fmt.Sprintf("(%s)", w)
		}
		tx = tx.Where(strings.Join(where, " AND "))
	}
	err := tx.Find(&staff).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return nil, response.GormErrorResponse(err, "An error occurred in the database")
	}

	for _, user := range staff {
		var contributionCount int64
		err := s.db.WithContext(ctx).Model(&models.Opportunity{}).Where("staff_id = ?", user.ID).Count(&contributionCount).Error
		if err != nil {
			s.logger.With(ctx).Error(err)
			return nil, response.GormErrorResponse(err, "An error occurred in the database")
		}

		var lastLoginAt string
		// err = s.db.WithContext(ctx).Model(&models.UserLogin{}).Where("user_id = ?", user.ID).Order("created_at DESC").Limit(1).Pluck("created_at", &lastLoginAt).Error
		// if err != nil {
		// 	s.logger.With(ctx).Error(err)
		// 	return nil, response.GormErrorResponse(err, "An error occurred in the database")
		// }

		staffReportData = append(staffReportData, StaffReportResponseData{
			User:              user,
			ContributionCount: contributionCount,
			LastLoginAt:       lastLoginAt,
		})
	}

	return staffReportData, response.ErrorResponse{}
}

func (s *service) StaffReportCount(ctx context.Context, params StaffReportQueryRequestParams) (int64, response.ErrorResponse) {
	if !policy.CanSeeStaffReport(ctx) {
		s.logger.With(ctx).Error("You do not have permission to see staff report")
		return 0, response.ErrorForbidden("You do not have permission to see the staff report")
	}

	var count int64

	tx := s.db.WithContext(ctx).Model(&models.User{}).Where("contributor_id IS NOT NULL")
	where := makeStaffReportFilters(params)
	if len(where) > 0 {
		// wrap any where item with () to avoid conflict with other where items
		for i, w := range where {
			where[i] = fmt.Sprintf("(%s)", w)
		}
		tx = tx.Where(strings.Join(where, " AND "))
	}
	err := tx.Count(&count).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return 0, response.GormErrorResponse(err, "An error occurred in the database")
	}

	return count, response.ErrorResponse{}
}
