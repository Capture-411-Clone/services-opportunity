package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/wishlist/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: wishlist
 * @apiPath: /wishlists
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: AddWishlistRequest
 * @apiResponseRef: AddWishlistResponse
 * @apiSummary: add wishlist
 * @apiDescription: add wishlist
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) add(ctx echo.Context) error {
	var input = &endpoints.AddWishlistRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	wishlist, err := r.service.Add(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(wishlist))
}

/*
 * @apiDefine: AddWishlistResponse
 */
type AddWishlistResponse struct {
	StatusCode int             `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.Wishlist `json:"data" openapi:"$ref:Wishlist;type:object"`
}
