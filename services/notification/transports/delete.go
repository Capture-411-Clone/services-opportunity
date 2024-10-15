package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/services/notification/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: notification
 * @apiPath: /notifications
 * @apiMethod: DELETE
 * @apiStatusCode: 200
 * @apiRequestRef: DeleteNotificationRequestBody
 * @apiResponseRef: DeleteNotificationResponse
 * @apiSummary: Delete Notification
 * @apiDescription: Delete Notification
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) delete(ctx echo.Context) error {
	var input = &endpoints.DeleteNotificationRequestBody{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "خطا در ساخت آدرس"))
	}

	deletedNotifications, err := r.service.Delete(ctx.Request().Context(), *input)

	result := DeleteNotificationResponseData{
		IDs: deletedNotifications,
	}

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusCreated, response.Success(result))
}

/*
 * @apiDefine: DeleteNotificationResponse
 */
type DeleteNotificationResponse struct {
	StatusCode int                            `json:"status_code" openapi:"example:200"`
	Data       DeleteNotificationResponseData `json:"data" openapi:"$ref:DeleteNotificationResponseData;type:object;"`
}

/*
 * @apiDefine: DeleteNotificationResponseData
 */
type DeleteNotificationResponseData struct {
	IDs string `json:"ids" openapi:"example:1,2,3;"`
}
