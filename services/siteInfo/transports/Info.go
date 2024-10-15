package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: siteInfo
 * @apiPath: /siteInfo/info
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiResponseRef: GetSiteInfoResponse
 * @apiSummary: get siteInfo
 * @apiDescription: get siteInfo
 */
func (r resource) get(ctx echo.Context) error {
	site, err := r.service.Info(ctx.Request().Context())
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusOK, response.Success(site))
}

/*
 * @apiDefine: GetSiteInfoResponse
 */
type GetSiteInfoResponse struct {
	StatusCode int             `json:"status_code" openapi:"example:200;type:integer"`
	Data       models.SiteInfo `json:"data" openapi:"$ref:SiteInfo;type:object"`
}
