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
 * @apiDefine: CreateMarketRequest
 */
type CreateMarketRequest struct {
	Name    string `json:"name" openapi:"example:Technology Department;type:string"`
	Acronym string `json:"acronym" openapi:"example:TDE;type:string"`
}

func (c *CreateMarketRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Create(ctx context.Context, input CreateMarketRequest) (models.Market, response.ErrorResponse) {
	if !policy.CanCreateMarket(ctx) {
		s.logger.With(ctx).Error("You do not have permission to create a market")
		return models.Market{}, response.ErrorForbidden("You do not have permission to create a market")
	}

	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	market := models.Market{
		UserID:  id,
		Name:    input.Name,
		Acronym: input.Acronym,
	}
	err := s.db.WithContext(ctx).Create(&market).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return market, response.GormErrorResponse(err, "An error occurred while creating the market")
	}
	return market, response.ErrorResponse{}
}
