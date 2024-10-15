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
)

/*
 * @apiDefine: FiscalYearQueryFilterType
 */
type FiscalYearQueryFilterType struct {
	Year        filter.FilterValue[string] `json:"year,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	SimilarName filter.FilterValue[string] `json:"similar_name" openapi:"nullable;$ref:FilterValueString;"`
	CreatedAt   filter.FilterValue[string] `json:"created_at" openapi:"nullable;$ref:FilterValueString;"`
}

/*
 * @apiDefine: FiscalYearQueryRequestParams
 */
type FiscalYearQueryRequestParams struct {
	ID      string                    `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page    string                    `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage string                    `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order   string                    `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy string                    `json:"order_by,omitempty" openapi:"example:year;in:query"`
	Query   string                    `json:"query,omitempty" openapi:"pattern:^[0-9]+$;in:query"`
	Filters FiscalYearQueryFilterType `json:"filters,omitempty" openapi:"$ref:FiscalYearQueryFilterType;in:query"`
}

func (data *FiscalYearQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":       gocriteria.New("id").Optional(),
		"query":    gocriteria.New("query").Optional(),
		"page":     gocriteria.New("page").Optional(),
		"per_page": gocriteria.New("per_page").Optional(),
		"order":    gocriteria.New("order").Optional(),
		"order_by": gocriteria.New("order_by").Optional(),
		"filters": gocriteria.Schema{
			"similar_name": gocriteria.Schema{
				"op":    gocriteria.New("filter.name.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.name.value").Optional(),
			},
			"year": gocriteria.Schema{
				"op":    gocriteria.New("filter.year.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.year.value").Optional(),
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

func makeFilters(params FiscalYearQueryRequestParams) []string {
	var where []string

	if params.ID != "" {
		where = append(where, fmt.Sprintf("id = %s", params.ID))
	}

	if params.Filters.SimilarName.Value != "" {
		opValue := filter.GetDBOperatorAndValue("equals", params.Filters.SimilarName.Value)
		where = append(where, fmt.Sprintf("LOWER(year) %% LOWER('%s')", opValue.Value))
	}

	if params.Filters.Year.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Year.Op, params.Filters.Year.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("year %s %s", opValue.Operator, val))
	}

	if params.Filters.CreatedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CreatedAt.Op, params.Filters.CreatedAt.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("created_at %s %s", opValue.Operator, val))
	}

	return where

}

func (s *service) Query(
	ctx context.Context, offset, limit int, params FiscalYearQueryRequestParams,
) ([]models.FiscalYear, response.ErrorResponse) {
	var fiscalYears []models.FiscalYear
	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order)).
		Preload("User")
	where := makeFilters(params)

	if len(where) > 0 {
		tx = tx.Where(strings.Join(where, " AND "))
	}

	err := tx.Find(&fiscalYears).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.FiscalYear{}, response.GormErrorResponse(err, "error in finding fiscal year")
	}

	return fiscalYears, response.ErrorResponse{}
}
