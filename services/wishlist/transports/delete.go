package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/wishlist/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: wishlist
 * @apiPath: /wishlists
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteWishlistRequest
 * @apiResponseRef: DeleteWishlistResponse
 * @apiSummary: delete wishlist
 * @apiDescription: delete wishlist
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteWishlistRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting wishlist"))
	}
	deletedWishlists, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedWishlists))
}

/*
 * @apiDefine: DeleteWishlistResponse
 */
type DeleteWishlistResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	Data       string `json:"data" openapi:"example:1,3,4;type:string"`
}
