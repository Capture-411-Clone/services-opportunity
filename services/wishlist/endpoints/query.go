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
 * @apiDefine: WishlistQueryFilterType
 */
type WishlistQueryFilterType struct {
	OpportunityID filter.FilterValue[int]    `json:"opportunity_id" openapi:"nullable;$ref:FilterValueInt;"`
	UserID        filter.FilterValue[int]    `json:"user_id" openapi:"nullable;$ref:FilterValueInt;"`
	CreatedAt     filter.FilterValue[string] `json:"created_at" openapi:"nullable;$ref:FilterValueString;"`
}

/*
 * @apiDefine: WishlistQueryRequestParams
 */
type WishlistQueryRequestParams struct {
	ID      string                  `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page    string                  `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage string                  `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order   string                  `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy string                  `json:"order_by,omitempty" openapi:"example:opportunity_id;in:query"`
	Query   string                  `json:"query,omitempty" openapi:"pattern:^[0-9]+$;in:query"`
	Filters WishlistQueryFilterType `json:"filters,omitempty" openapi:"$ref:WishlistQueryFilterType;in:query"`
}

func (data *WishlistQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":       gocriteria.New("id").Optional(),
		"page":     gocriteria.New("page").Optional(),
		"per_page": gocriteria.New("per_page").Optional(),
		"query":    gocriteria.New("query").Optional(),
		"order":    gocriteria.New("order").Optional(),
		"order_by": gocriteria.New("order_by").Optional(),
		"filters": gocriteria.Schema{
			"opportunity_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.opportunity_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.opportunity_id.value").Optional(),
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

func makeFilters(params WishlistQueryRequestParams) []string {
	var where []string

	if params.ID != "" {
		where = append(where, fmt.Sprintf("id = %s", params.ID))
	}

	if params.Filters.OpportunityID.Op != "" {
		OpportunityIDStr := strconv.Itoa(params.Filters.OpportunityID.Value)
		opValue := filter.GetDBOperatorAndValue(params.Filters.OpportunityID.Op, OpportunityIDStr)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("opportunity_id %s %s", opValue.Operator, val))
	}
	if params.Filters.UserID.Op != "" {
		UserIDStr := strconv.Itoa(params.Filters.UserID.Value)
		opValue := filter.GetDBOperatorAndValue(params.Filters.UserID.Op, UserIDStr)
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

func (s *service) Query(
	ctx context.Context, offset, limit int, params WishlistQueryRequestParams,
) ([]models.Wishlist, response.ErrorResponse) {
	var wishlists []models.Wishlist
	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order)).
		Preload("User").
		Preload("Opportunity").
		Preload("Opportunity.Documents")
	where := makeFilters(params)

	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	if !policy.IsAdmin(ctx) || !policy.IsReviewer(ctx) {
		tx = tx.Where("user_id = ?", policy.ExtractIdClaim(ctx))
	}

	err := tx.Find(&wishlists).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Wishlist{}, response.GormErrorResponse(err, "error in finding fiscal year")
	}

	return wishlists, response.ErrorResponse{}
}
