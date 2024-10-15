package endpoints

import (
	"context"
	"strconv"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/product"
)

func (s *service) GetAllProducts(ctx context.Context) ([]stripe.Product, response.ErrorResponse) {
	var captureProducts []stripe.Product
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.First(&user, theID).Error
	if err != nil {
		return []stripe.Product{}, response.ErrorNotFound("User not found")
	}

	// List all products
	params := &stripe.ProductListParams{}
	// params.AddExpand("data.default_price") // Optionally expand the default price

	i := product.List(params)

	for i.Next() {
		prod := i.Product()
		if app, ok := prod.Metadata["app"]; ok && app == "capture411" {
			captureProducts = append(captureProducts, *prod)
		}
	}

	return captureProducts, response.ErrorResponse{}
}
