package transports

import (
	"net/http"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/notification/endpoints"
	"github.com/labstack/echo/v4"
)

/*
 * @apiTag: notification
 * @apiPath: /notifications/{id}
 * @apiMethod: PATCH
 * @apiStatusCode: 200
 * @apiParametersRef: UpdateNotificationRequestParams
 * @apiRequestRef: UpdateNotificationRequestBody
 * @apiResponseRef: UpdateNotificationResponse
 * @apiSummary: Update notification
 * @apiDescription: Update notification
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) update(ctx echo.Context) error {
	params := &endpoints.UpdateNotificationRequestParams{}
	p := map[string]string{
		"id": ctx.Param("id"),
	}
	perrors := params.Validate(p)
	if perrors != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(perrors))
	}

	var input = &endpoints.UpdateNotificationRequestBody{}
	errors := input.Validate(ctx.Request())
	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors))
	}

	updatedNotification, err := r.service.Update(ctx.Request().Context(), params.ID, *input)

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}

	return ctx.JSON(http.StatusOK, response.Success(updatedNotification))
}

/*
 * @apiDefine: UpdateNotificationResponse
 */
type UpdateNotificationResponse struct {
	StatusCode int                 `json:"status_code"`
	Data       models.Notification `json:"data" openapi:"$ref:Notification;type:object;"`
}
