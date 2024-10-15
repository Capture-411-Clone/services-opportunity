package endpoints

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: CreateFiscalYearRequest
 */
type CreateFiscalYearRequest struct {
	Year string `json:"year" openapi:"example:2022;type:string;"`
}

func (c *CreateFiscalYearRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Create(ctx context.Context, input CreateFiscalYearRequest) (models.FiscalYear, response.ErrorResponse) {
	if !policy.CanCreateFiscalYear(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a fiscalYear")
		return models.FiscalYear{}, response.ErrorForbidden("You do not have permission to create a fiscalYear")
	}

	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	fiscalYear := models.FiscalYear{
		UserID: id,
		Year:   input.Year,
	}
	err := s.db.WithContext(ctx).Create(&fiscalYear).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return fiscalYear, response.GormErrorResponse(err, "An error occurred while creating the fiscalYear")
	}
	return fiscalYear, response.ErrorResponse{}
}
