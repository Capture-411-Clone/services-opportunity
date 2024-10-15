package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/siteInfo/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: siteInfo
 * @apiPath: /siteInfo/update/{id}
 * @apiMethod: PUT
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateSiteInfoParams
 * @apiRequestRef: UpdateSiteInfoRequest
 * @apiResponseRef: UpdateSiteInfoResponse
 * @apiSummary: update siteInfo
 * @apiDescription: update siteInfo
 * @apiSecurity: jwtBearerAuth
 */
func (r *resource) update(ctx echo.Context) error {
	var input = &endpoints.UpdateSiteInfoRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	site, err := r.service.Update(ctx.Request().Context(), *input)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Success(site))
}

/*
 * @apiDefine: UpdateSiteInfoResponse
 */
type UpdateSiteInfoResponse struct {
	StatusCode int             `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.SiteInfo `json:"data" openapi:"$ref:SiteInfo;type:object"`
}
