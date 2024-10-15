package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/contractVehicle/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: contractVehicle
 * @apiPath: /contractVehicles/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: ContractVehicleUpdateParams
 * @apiRequestRef: UpdateContractVehicleRequest
 * @apiResponseRef: UpdateContractVehicleResponse
 * @apiSummary: update contractVehicle
 * @apiDescription: update contractVehicle
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	params := &endpoints.ContractVehicleUpdateParams{}
	p := gocriteria.Params{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}
	var input = &endpoints.UpdateContractVehicleRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	contractVehicle, err := r.service.Update(ctx.Request().Context(), params.ID, *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(contractVehicle))
}

/*
 * @apiDefine: UpdateContractVehicleResponse
 */
type UpdateContractVehicleResponse struct {
	StatusCode int                      `json:"status_code" openapi:"example:200;type:integer"`
	Data       []models.ContractVehicle `json:"data" openapi:"$ref:ContractVehicle;type:object"`
}
