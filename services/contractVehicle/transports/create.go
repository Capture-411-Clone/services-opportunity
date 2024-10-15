package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/contractVehicle/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: contractVehicle
 * @apiPath: /contractVehicles
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateContractVehicleRequest
 * @apiResponseRef: CreateContractVehicleResponse
 * @apiSummary: create contractVehicle
 * @apiDescription: create contractVehicle
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateContractVehicleRequest{}

	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	contractVehicle, err := r.service.Create(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(contractVehicle))
}

/*
 * @apiDefine: CreateContractVehicleResponse
 */
type CreateContractVehicleResponse struct {
	StatusCode int                      `json:"status_code" openapi:"example:200;type:integer"`
	Data       []models.ContractVehicle `json:"data" openapi:"$ref:ContractVehicle;type:object"`
}
