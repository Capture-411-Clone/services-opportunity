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
 * @apiDefine: AgencyQueryFilterType
 */
type AgencyQueryFilterType struct {
	Name         filter.FilterValue[string] `json:"name" openapi:"nullable;$ref:FilterValueString;"`
	SimilarName  filter.FilterValue[string] `json:"similar_name" openapi:"nullable;$ref:FilterValueString;"`
	Acronym      filter.FilterValue[string] `json:"acronym" openapi:"nullable;$ref:FilterValueString;example:{\"title\":{\"op\":\"contains\",\"value\":\"at\"}"`
	DepartmentID filter.FilterValue[int]    `json:"department_id" openapi:"nullable;$ref:FilterValueInt;"`
	CreatedAt    filter.FilterValue[string] `json:"created_at" openapi:"nullable;$ref:FilterValueString;"`
}

/*
 * @apiDefine: AgencyQueryRequestParams
 */
type AgencyQueryRequestParams struct {
	ID      string                `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page    string                `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage string                `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order   string                `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy string                `json:"order_by,omitempty" openapi:"example:name;in:query"`
	Query   string                `json:"query,omitempty" openapi:"pattern:^[0-9]+$;in:query"`
	Filters AgencyQueryFilterType `json:"filters,omitempty" openapi:"$ref:AgencyQueryFilterType;in:query"`
}

func (data *AgencyQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":               gocriteria.New("id").Optional(),
		"page":             gocriteria.New("page").Optional(),
		"per_page":         gocriteria.New("per_page").Optional(),
		"query":            gocriteria.New("query").Optional(),
		"order":            gocriteria.New("order").Optional(),
		"order_by":         gocriteria.New("order_by").Optional(),
		"full_text_search": gocriteria.New("full_text_search").Optional(),
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
			"department_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.department_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.department_id.value").Optional(),
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

func makeFilters(params AgencyQueryRequestParams) []string {
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

	if params.Filters.DepartmentID.Op != "" {
		DepartmentIDStr := strconv.Itoa(params.Filters.DepartmentID.Value)
		opValue := filter.GetDBOperatorAndValue(params.Filters.DepartmentID.Op, DepartmentIDStr)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("department_id %s %s", opValue.Operator, val))
	}

	if params.Filters.CreatedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CreatedAt.Op, params.Filters.CreatedAt.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("created_at %s %s", opValue.Operator, val))
	}

	return where

}

func query(params AgencyQueryRequestParams) []string {
	var where []string
	if params.Query != "" {
		fields := []string{
			"agencies.name",
			"agencies.acronym",
			"departments.name",
		}

		for _, field := range fields {
			where = append(where, fmt.Sprintf("%s ILIKE %s", field, fmt.Sprintf("'%%%s%%'", strings.TrimSpace(params.Query))))
		}
	}
	return where
}

func (s *service) Query(
	ctx context.Context, offset, limit int, params AgencyQueryRequestParams,
) ([]models.Agency, response.ErrorResponse) {
	var agencies []models.Agency
	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order)).
		Preload("User").Preload("Offices").Preload("Department")
	if params.Query != "" {
		tx = tx.
			Joins("LEFT JOIN departments ON agencies.department_id = departments.id")
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

	err := tx.Find(&agencies).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Agency{}, response.GormErrorResponse(err, "error in finding agency")
	}

	return agencies, response.ErrorResponse{}
}
