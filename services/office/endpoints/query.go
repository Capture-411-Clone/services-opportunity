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
* @apiDefine: OfficeQueryFilterType
 */
type OfficeQueryFilterType struct {
	Name        filter.FilterValue[string] `json:"name,omitempty" openapi:"nullable;$ref:FilterValueString;"`
	SimilarName filter.FilterValue[string] `json:"similar_name" openapi:"nullable;$ref:FilterValueString;"`
	Acronym     filter.FilterValue[string] `json:"acronym" openapi:"nullable;$ref:FilterValueString;;in:path"`
	ParentID    filter.FilterValue[int]    `json:"parent_id" openapi:"nullable;$ref:FilterValueInt;"`
	AgencyID    filter.FilterValue[int]    `json:"agency_id" openapi:"nullable;$ref:FilterValueInt;"`
	CreatedAt   filter.FilterValue[string] `json:"created_at" openapi:"nullable;$ref:FilterValueString;"`
}

/*
* @apiDefine: OfficeQueryRequestParams
 */
type OfficeQueryRequestParams struct {
	ID      string                `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	Page    string                `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage string                `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order   string                `json:"order,omitempty" openapi:"example:desc;in:query"`
	OrderBy string                `json:"order_by,omitempty" openapi:"example:name;in:query"`
	Query   string                `json:"query,omitempty" openapi:"pattern:^[0-9]+$;in:query"`
	Filters OfficeQueryFilterType `json:"filters,omitempty" openapi:"$ref:OfficeQueryFilterType;in:query"`
}

func (data *OfficeQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
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
			"parent_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.parent_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.parent_id.value").Optional(),
			},
			"agency_id": gocriteria.Schema{
				"op":    gocriteria.New("filter.agency_id.op").FilterOperators().Optional(),
				"value": gocriteria.New("filter.agency_id.value").Optional(),
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

func makeFilters(params OfficeQueryRequestParams) []string {
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
		where = append(where, fmt.Sprintf("(ts_rank_cd(to_tsvector('english', LOWER(name)), %s) + similarity(LOWER(name), LOWER('%s'))) / 2 > 0.1", tsQuery, normalizedInput))
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

	if params.Filters.ParentID.Op != "" {
		ParentIDStr := strconv.Itoa(params.Filters.ParentID.Value)
		opValue := filter.GetDBOperatorAndValue(params.Filters.ParentID.Op, ParentIDStr)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("parent_id %s %s", opValue.Operator, val))
	}

	if params.Filters.AgencyID.Op != "" {
		agencyIDStr := strconv.Itoa(params.Filters.AgencyID.Value)
		opValue := filter.GetDBOperatorAndValue(params.Filters.AgencyID.Op, agencyIDStr)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("agency_id %s %s", opValue.Operator, val))
	}

	if params.Filters.CreatedAt.Op != "" {
		opValue := filter.GetDBOperatorAndValue(params.Filters.CreatedAt.Op, params.Filters.CreatedAt.Value)
		val := exp.TerIf(opValue.Value == "", "", fmt.Sprintf("'%s'", opValue.Value))
		where = append(where, fmt.Sprintf("created_at %s %s", opValue.Operator, val))
	}

	return where

}

func query(params OfficeQueryRequestParams) []string {
	var where []string
	if params.Query != "" {
		fields := []string{
			"offices.name",
			"offices.acronym",
			"agencies.name",
		}
		for _, field := range fields {
			where = append(where, fmt.Sprintf("%s ILIKE %s", field, fmt.Sprintf("'%%%s%%'", strings.TrimSpace(params.Query))))
		}
	}
	return where
}

func (s *service) Query(
	ctx context.Context, offset, limit int, params OfficeQueryRequestParams,
) ([]models.Office, response.ErrorResponse) {
	var offices []models.Office
	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order)).
		Preload("User").
		Preload("Agency").
		Preload("Parent.Parent.Parent").
		Preload("Children.Children.Children")

	if params.Query != "" {
		tx = tx.
			Joins("LEFT JOIN agencies ON offices.agency_id = agencies.id")
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

	err := tx.Find(&offices).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Office{}, response.GormErrorResponse(err, "error in finding office")
	}

	return offices, response.ErrorResponse{}
}
