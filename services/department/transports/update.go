package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/department/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: department
 * @apiPath: /departments/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateDepartmentRequestParam
 * @apiRequestRef: UpdateDepartmentRequest
 * @apiResponseRef: UpdateDepartmentResponse
 * @apiSummary: update department
 * @apiDescription: update department
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateDepartmentRequestParam{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateDepartmentRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	department, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(department))
}

/*
 * @apiDefine: UpdateDepartmentResponse
 */
type UpdateDepartmentResponse struct {
	StatusCode int                 `json:"status_code" openapi:"example:200;type:integer"`
	Data       []models.Department `json:"data" openapi:"$ref:Department;type:array"`
}
