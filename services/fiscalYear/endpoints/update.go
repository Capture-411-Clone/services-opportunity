package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateFiscalYearRequest
 */
type UpdateFiscalYearRequest struct {
	Year string `json:"year" openapi:"example:2022;type:string;"`
}

func (c *UpdateFiscalYearRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"year": gocriteria.New("year").Required().MinMaxLength(2, 200),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"year": "Years",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

/*
 * @apiDefine: UpdateFiscalYearRequestParam
 */
type UpdateFiscalYearRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateFiscalYearRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Update(ctx context.Context, id string, input UpdateFiscalYearRequest) (
	models.FiscalYear, response.ErrorResponse,
) {
	var fiscalYear models.FiscalYear

	if !policy.CanUpdateFiscalYear(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a fiscalYear")
		return models.FiscalYear{}, response.ErrorForbidden(nil, "You do not have permission to update a fiscalYear")
	}

	err := s.db.WithContext(ctx).First(&fiscalYear, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.FiscalYear{}, response.GormErrorResponse(err, "An error occurred while finding the fiscalYear")
	}
	fiscalYear.Year = input.Year

	err = s.db.WithContext(ctx).Save(&fiscalYear).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.FiscalYear{}, response.GormErrorResponse(err, "An error occurred while updating the fiscalYear")
	}
	return fiscalYear, response.ErrorResponse{}
}
