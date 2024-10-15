package endpoints

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/modules/calculator"
)

/*
 * @apiDefine: CalculateOpportunityPriceRequestParam
 */
type CalculateOpportunityPriceRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *CalculateOpportunityPriceRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "Opportunity ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) CalculateOpportunityPrice(ctx context.Context, id string) (float64, response.ErrorResponse) {

	var opportunity models.Opportunity

	err := s.db.Model(models.Opportunity{}).Where("id = ?", id).First(&opportunity).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Calc Opportunity Price: %v", err)
		return 0, response.GormErrorResponse(err, "error calculating opportunity price")
	}

	contractValueFloat64 := 0.0
	if opportunity.CrawlerContractValue != "" {
		contractValueFloat64, err = strconv.ParseFloat(opportunity.CrawlerContractValue, 64)
		if err != nil {
			s.logger.With(ctx).Errorf("CreateOpportunityOrderCheckout: %v", err)
			return 0, response.ErrorBadRequest(nil, "Could not proceed with payment process")
		}
	}

	userContractValueFloat64 := 0.0
	if opportunity.UserContractValue != "" {
		userContractValueFloat64, err = strconv.ParseFloat(opportunity.UserContractValue, 64)
		if err != nil {
			s.logger.With(ctx).Errorf("CreateOpportunityOrderCheckout: %v", err)
			return 0, response.ErrorBadRequest(nil, "Could not proceed with payment process")
		}
	}

	// Decide which value to use
	var finalValue float64
	if userContractValueFloat64 != 0.0 {
		finalValue = userContractValueFloat64
	} else {
		finalValue = contractValueFloat64
	}

	finalPrice := calculator.CalculateOpportunityPrice(finalValue)

	return finalPrice, response.ErrorResponse{}
}
