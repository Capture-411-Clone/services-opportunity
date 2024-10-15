package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/pagination"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/notification/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: notification
 * @apiMethod: GET
 * @apiStatusCode: 200
 * @apiDeprecated: false
 * @apiPath: /notifications
 * @apiParametersRef: NotificationQueryRequestParams
 * @apiResponseRef: NotificationQueryResponse
 * @apiSummary: Get notifications list
 * @apiErrorStatusCodes: 400, 404
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) query(ctx echo.Context) error {
	var params = &endpoints.NotificationQueryRequestParams{}

	errs := params.ValidateQueries(ctx.Request())
	if len(errs) > 0 {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errs, "query parameter is not valid"))
	}

	if params.OrderBy == "" {
		params.OrderBy = "id"
	}
	if params.Order == "" {
		params.Order = "desc"
	}

	c := ctx.Request().Context()

	count, err := r.service.Count(c, *params)
	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	pages := pagination.NewFromRequest(ctx.Request(), int(count))
	notifications, er := r.service.Query(
		c, pages.Offset(), pages.Limit(),
		*params,
	)
	if er.StatusCode != 0 {
		return ctx.JSON(er.StatusCode, er)
	}

	result := NotificationQueryResponseData{
		Limit:      pages.Limit(),
		Offset:     pages.Offset(),
		Page:       pages.Page,
		TotalRows:  int64(pages.TotalCount),
		TotalPages: pages.PageCount,
		Items:      notifications,
	}

	return ctx.JSON(http.StatusOK, response.Success(result))
}

/*
 * @apiDefine: NotificationQueryResponse
 */
type NotificationQueryResponse struct {
	StatusCode int                           `json:"status_code" openapi:"example:200;type:integer"`
	Data       NotificationQueryResponseData `json:"data" openapi:"$ref:NotificationQueryResponseData;type:object;"`
}

/*
 * @apiDefine: NotificationQueryResponseData
 */
type NotificationQueryResponseData struct {
	Limit      int                   `json:"limit" openapi:"example:10;type:integer;"`
	Offset     int                   `json:"offset" openapi:"example:0;type:integer;"`
	Page       int                   `json:"page" openapi:"example:1;type:integer;"`
	TotalRows  int64                 `json:"totalRows" openapi:"example:20;type:integer;"`
	TotalPages int                   `json:"totalPages" openapi:"example:2;type:integer;"`
	Items      []models.Notification `json:"items" openapi:"$ref:Notification;type:array;"`
}
