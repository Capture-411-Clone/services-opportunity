package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/exp"
	"github.com/Capture-411/services-opportunity/kit/filter"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

/*
 * @apiDefine: DepartmentQueryFilterType
 */
type DepartmentQueryFilterType struct {
	Name        filter.FilterValue[string] `json:"name" openapi:"nullable;$ref:FilterValueString;"`
	SimilarName filter.FilterValue[string] `json:"similar_name" openapi:"nullable;$ref:FilterValueString;"`
	MarketID    filter.FilterValue[int]    `json:"market_id" openapi:"nullable;$ref:FilterValueInt;"`
	Acronym     filter.FilterValue[string] `json:"acronym" openapi:"nullable;$ref:FilterValueString;example:{\"title\":{\"op\":\"contains\",\"value\":\"at\"}"`
	CreatedAt   filter.FilterValue[string] `json:"created_at" openapi:"nullable;$ref:FilterValueString;"`
}

/*
 * @apiDefine: DepartmentQueryRequestParams
 */
type DepartmentQueryRequestParams struct {
	ID      string                    `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page    string                    `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage string                    `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order   string                    `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy string                    `json:"order_by,omitempty" openapi:"example:name;in:query"`
	Query   string                    `json:"query,omitempty" openapi:"pattern:^[0-9]+$;in:query"`
	Filters DepartmentQueryFilterType `json:"filters,omitempty" openapi:"$ref:DepartmentQueryFilterType;in:query"`
}

func (data *DepartmentQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {

	schema := gocriteria.Schema{
		"id":       gocriteria.New("id").Optional(),
		"page":     gocriteria.New("page").Optional(),
		"per_page": gocriteria.New("per_page").Optional(),
		"query":    gocriteria.New("query").Optional(),
		"order":    gocriteria.New("order").Optional(),
		"order_by": gocriteria.New("order_by").Optional(),
		"filters": gocriteria.Schema{
			"similar_name": gocriteria.Schema{
				"op":    gocriteria.New("filter.name.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.name.value").Optional(),
			},
			"name": gocriteria.Schema{
				"op":    gocriteria.New("filter.name.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.name.value").Optional(),
			},
			"acronym": gocriteria.Schema{
				"op":    gocriteria.New("filter.acronym.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.acronym.value").Optional(),
			},
			"market_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.market_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.market_id.value").Optional(),
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

func makeFilters(params DepartmentQueryRequestParams) []string {
	var where []string

	if params.ID != "" {
		where = append(where, fmt.Sprintf("id = %s", params.ID))
	}

	if params.Filters.SimilarName.Value != "" {
		normalizedInput := strings.ToLower(params.Filters.SimilarName.Value)
		// remove specia chars
		rgx := regexp.MustCompile("[^a-zA-Z0-9]+")
		normalizedInput = rgx.ReplaceAllString(normalizedInput, " ")

		// Prepare the full-text search vector and query. Adjust the configuration ('english') as needed.
		tsQuery := fmt.Sprintf("plainto_tsquery('english', LOWER('%s'))", normalizedInput)
		// Combine full-text search with trigram similarity for ranking
		where = append(where, fmt.Sprintf("(ts_rank_cd(to_tsvector('english', LOWER(name)), %s) + similarity(LOWER(name), LOWER('%s'))) / 2 > 0.2", tsQuery, normalizedInput))
		// This assumes an additive model for ranking combining both scores, adjust the threshold and weights as needed.
	}

	if params.Filters.Name.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Name.Op, params.Filters.Name.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("name %s %s", opValue.Operator, val))
	}

	if params.Filters.Acronym.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.Acronym.Op, params.Filters.Acronym.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("acronym %s %s", opValue.Operator, val))
	}

	if params.Filters.MarketID.Op != "" {
		MarketIDStr := strconv.Itoa(params.Filters.MarketID.Value)
		opValue := filter.GetDBOperatorAndValue(params.Filters.MarketID.Op, MarketIDStr)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("market_id %s %s", opValue.Operator, val))
	}

	if params.Filters.CreatedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CreatedAt.Op, params.Filters.CreatedAt.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("created_at %s %s", opValue.Operator, val))
	}

	return where

}

func query(params DepartmentQueryRequestParams) []string {
	var where []string
	if params.Query != "" {
		fields := []string{
			"departments.name",
			"departments.acronym",
			"markets.name",
		}
		for _, field := range fields {
			where = append(where, fmt.Sprintf("%s ILIKE %s", field, fmt.Sprintf("'%%%s%%'", strings.TrimSpace(params.Query))))
		}
	}
	return where
}

func (s *service) Query(
	ctx context.Context, offset, limit int, params DepartmentQueryRequestParams,
) ([]models.Department, response.ErrorResponse) {
	var departments []models.Department
	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order)).
		Preload("User").Preload("Agencies").Preload("Market")
	if params.Query != "" {
		tx = tx.
			Joins("LEFT JOIN markets ON departments.market_id = markets.id")
		where := query(params)
		for i, w := range where {
			where[i] = fmt.Sprintf("(%s)", w)
		}
		tx = tx.Where(strings.Join(where, " OR "))
	} else {
		where := makeFilters(params)
		for i, w := range where {
			where[i] = fmt.Sprintf("(%s)", w)
		}
		if len(where) > 0 {
			tx = tx.Where(strings.Join(where, " AND "))
		}
	}
	err := tx.Debug().Find(&departments).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Department{}, response.GormErrorResponse(err, "error in finding agency")
	}

	return departments, response.ErrorResponse{}
}
