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
 * @apiDefine: UpdateMarketRequest
 */
type UpdateMarketRequest struct {
	Name    string `json:"name" openapi:"example:Technology Department;type:string"`
	Acronym string `json:"acronym" openapi:"example:TDE;type:string"`
}

func (c *UpdateMarketRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"name":    gocriteria.New("name").Required().MinMaxLength(2, 200),
		"acronym": gocriteria.New("acronym"),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"name":    "Name",
			"acronym": "Acronym",
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
 * @apiDefine: UpdateMarketRequestParam
 */
type UpdateMarketRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateMarketRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
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

func (s *service) Update(ctx context.Context, id string, input UpdateMarketRequest) (
	models.Market, response.ErrorResponse,
) {
	var market models.Market

	if !policy.CanUpdateMarket(ctx) {
		s.logger.With(ctx).Error("You do not have permission to update a market")
		return models.Market{}, response.ErrorForbidden(nil, "You do not have permission to update a market")
	}

	err := s.db.WithContext(ctx).First(&market, "id", id).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Market{}, response.GormErrorResponse(err, "An error occurred while finding the market")
	}
	market.Name = input.Name
	market.Acronym = input.Acronym

	err = s.db.WithContext(ctx).Save(&market).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Market{}, response.GormErrorResponse(err, "An error occurred while updating the market")
	}
	return market, response.ErrorResponse{}
}
