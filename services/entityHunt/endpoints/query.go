package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: EntityHuntQueryRequestParams
 */
type EntityHuntQueryRequestParams struct {
	ID                 string `json:"id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	LastVisitedID      string `json:"last_visited_id,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	SolicitationNumber string `json:"solicitation_number,omitempty" openapi:"example:123456;pattern:^[0-9]+$;in:query"`
	Year               string `json:"year,omitempty" openapi:"example:2022;pattern:^[0-9]+$;in:query"`
	FileName           string `json:"file_name,omitempty" openapi:"example:filename;in:query"`
	FilePath           string `json:"file_path,omitempty" openapi:"example:user/documents/report.pdf;in:query"`
	Page               string `json:"page,omitempty" openapi:"example:1;pattern:^[0-9]+$;in:query"`
	PerPage            string `json:"per_page,omitempty" openapi:"example:10;pattern:^[0-9]+$;in:query"`
	Order              string `json:"order,omitempty" openapi:"example:asc;in:query"`
	OrderBy            string `json:"order_by,omitempty" openapi:"example:id;in:query"`
	Query              string `json:"query,omitempty" openapi:"pattern:^[0-9]+$;in:query"`
}

func (data *EntityHuntQueryRequestParams) ValidateQueries(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id":                  gocriteria.New("id").Optional(),
		"last_visited_id":     gocriteria.New("last_visited_id").Optional(),
		"solicitation_number": gocriteria.New("solicitation_number").Optional(),
		"year":                gocriteria.New("year").Optional(),
		"file_name":           gocriteria.New("file_name").Optional(),
		"file_path":           gocriteria.New("file_path").Optional(),
		"page":                gocriteria.New("page").Optional(),
		"per_page":            gocriteria.New("per_page").Optional(),
		"query":               gocriteria.New("query").Optional(),
		"order":               gocriteria.New("order").Optional(),
		"order_by":            gocriteria.New("order_by").Optional(),
	}
	errs := gocriteria.ValidateQueries(r, schema, data)
	if len(errs) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errs)
		return dumpedErrors
	}

	return nil
}

func (s *service) Query(
	ctx context.Context, offset, limit int, params EntityHuntQueryRequestParams,
) ([]models.EntityHunt, response.ErrorResponse) {

	if !policy.CanQueryEntityHunt(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a entityHunt")
		return []models.EntityHunt{}, response.ErrorForbidden("You do not have permission to create a entityHunt")
	}
	var entityHunts []models.EntityHunt
	tx := s.db.WithContext(ctx).Offset(offset).Limit(limit).
		Order(fmt.Sprintf("%s %s", params.OrderBy, params.Order))

	if params.LastVisitedID != "" {
		tx = tx.Where("id > ?", params.LastVisitedID)
	}

	if params.SolicitationNumber != "" {
		tx = tx.Where("solicitation_number = ?", params.SolicitationNumber)
	}

	if params.Year != "" {
		tx = tx.Where("year = ?", params.Year)
	}

	if params.FileName != "" {
		tx = tx.Where("file_name = ?", params.FileName)
	}

	if params.FilePath != "" {
		tx = tx.Where("file_path = ?", params.FilePath)
	}

	err := tx.Find(&entityHunts).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.EntityHunt{}, response.GormErrorResponse(err, "error in finding entityHunt")
	}

	return entityHunts, response.ErrorResponse{}
}
