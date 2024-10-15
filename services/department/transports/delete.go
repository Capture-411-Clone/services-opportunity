package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/department/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: department
 * @apiPath: /departments
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteDepartmentRequest
 * @apiResponseRef: DeleteDepartmentResponse
 * @apiSummary: delete department
 * @apiDescription: delete department
 * @apiSecurity: jwtBearerAuth
 */

func (r *resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteDepartmentRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "error in deleting agency"))
	}
	deletedCategories, err := r.service.Delete(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(deletedCategories))
}

/*
 * @apiDefine: DeleteDepartmentResponse
 */
type DeleteDepartmentResponse struct {
	StatusCode int    `json:"status_code" openapi:"example:200;type:integer"`
	IDs        string `json:"data" openapi:"example:1,2,3;"`
}
