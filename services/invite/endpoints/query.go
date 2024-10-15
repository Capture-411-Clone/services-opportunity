package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/exp"
	"github.com/Capture-411/services-opportunity/kit/filter"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: InviteQueryFilterType
 */
type InviteQueryFilterType struct {
	Code      filter.FilterValue[string] `json:"code,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	Limit     filter.FilterValue[int]    `json:"limit,omitempty" openapi:"nullable;$ref:FilterValueInt;"`
	UserID    filter.FilterValue[int]    `json:"user_id,omitempty" openapi:"nullable;$ref:FilterValueInt;"`
	CreatedAt filter.FilterValue[string] `json:"created_at" openapi:"nullable;$ref:FilterValueString;"`
}

/*
 * @apiDefine: InviteQueryRequestParams
 */
type InviteQueryRequestParams struct {
	ID      string                `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page    string                `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage string                `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order   string                `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy string                `json:"order_by,omitempty" openapi:"example:code;in:query"`
	Query   string                `json:"query,omitempty" openapi:"pattern:^[0-9]+$;in:query"`
	Filters InviteQueryFilterType `json:"filters,omitempty" openapi:"$ref:InviteQueryFilterType;in:query"`
}

func (data *InviteQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":       gocriteria.New("id").Optional(),
		"query":    gocriteria.New("query").Optional(),
		"page":     gocriteria.New("page").Optional(),
		"per_page": gocriteria.New("per_page").Optional(),
		"order":    gocriteria.New("order").Optional(),
		"order_by": gocriteria.New("order_by").Optional(),
		"filters": gocriteria.Schema{
			"code": gocriteria.Schema{
				"op":    gocriteria.New("filter.code.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.code.value").Optional(),
			},
			"limit": gocriteria.Schema{
				"op":    gocriteria.New("filter.limit.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.limit.value").Optional(),
			},
			"user_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.user_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.user_id.value").Optional(),
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

func makeFilters(params InviteQueryRequestParams) []string {
	var where []string

	if params.ID != "" {
		where = append(where, fmt.Sprintf("id = %s", params.ID))
	}

	if params.Filters.Code.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Code.Op, params.Filters.Code.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("code %s %s", opValue.Operator, val))
	}
	if params.Filters.Limit.Op != "" {
		limitStr := strconv.Itoa(params.Filters.Limit.Value)
		opValue := filter.GetDBOperatorAndValue(params.Filters.Limit.Op, limitStr)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("limit %s %s", opValue.Operator, val))
	}
	if params.Filters.UserID.Op != "" {
		usrIDStr := strconv.Itoa(params.Filters.UserID.Value)
		opValue := filter.GetDBOperatorAndValue(params.Filters.UserID.Op, usrIDStr)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("user_id %s %s", opValue.Operator, val))
	}

	if params.Filters.CreatedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CreatedAt.Op, params.Filters.CreatedAt.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("created_at %s %s", opValue.Operator, val))
	}

	return where

}
func (s *service) Query(ctx context.Context, offset, limit int, params InviteQueryRequestParams) (
	[]models.Invite, response.ErrorResponse,
) {
	var invites []models.Invite
	if !policy.CanQueryInvite(ctx) {
		s.logger.With(ctx).Error("you don't have permission to query invite")
		return invites, response.ErrorForbidden("you don't have permission to query invite")
	}
	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order)).
		Preload("User")
	where := makeFilters(params)

	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	err := tx.Find(&invites).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Invite{}, response.GormErrorResponse(err, "error in finding invite")
	}

	return invites, response.ErrorResponse{}
}
