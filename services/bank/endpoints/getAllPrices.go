package endpoints

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
)

type GetAllPricesRequestParams struct {
	ProductID string `json:"product_id" openapi:"example:prod_PbRzzvjoOCsfMQ;nullable;in:path"`
}

func (p *GetAllPricesRequestParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"product_id": gocriteria.New("product_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"product_id": "Product ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) GetAllPrices(ctx context.Context, params GetAllPricesRequestParams) ([]stripe.Price, response.ErrorResponse) {
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.First(&user, theID).Error
	if err != nil {
		return []stripe.Price{}, response.ErrorNotFound("User not found")
	}

	// For each product, fetch all associated prices
	priceParams := &stripe.PriceListParams{Product: stripe.String(params.ProductID)}
	priceList := price.List(priceParams)

	var prices []stripe.Price
	for priceList.Next() {
		p := priceList.Price()
		prices = append(prices, *p)
	}

	return prices, response.ErrorResponse{}
}
