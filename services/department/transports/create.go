package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/department/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: department
 * @apiPath: /departments
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateDepartmentRequest
 * @apiResponseRef: CreateDepartmentResponse
 * @apiSummary: create department
 * @apiDescription: create department
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateDepartmentRequest{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	department, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(department))
}

/*
 * @apiDefine: CreateDepartmentResponse
 */
type CreateDepartmentResponse struct {
	StatusCode int                 `json:"status_code" openapi:"example:200;type:integer"`
	Data       []models.Department `json:"data" openapi:"$ref:Department;type:array"`
}
