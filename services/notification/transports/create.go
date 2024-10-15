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
 * @apiPath: /notifications
 * @apiMethod: POST
 * @apiStatusCode: 201
 * @apiRequestRef: CreateNotificationRequest
 * @apiResponseRef: CreateNotificationResponse
 * @apiSummary: Create Notification
 * @apiDescription: Create Notification
 * @apiSecurity: jwtBearerAuth
 */
func (r resource) create(ctx echo.Context) error {
	var input = &endpoints.CreateNotificationRequest{}
	errors := input.Validate(ctx.Request())

	if errors != nil {
		r.logger.With(ctx.Request().Context()).Error(errors)
		return ctx.JSON(http.StatusBadRequest, response.ErrorBadRequest(errors, "input data is invalid"))
	}

	notification, err := r.service.Create(ctx.Request().Context(), *input)

	if err.StatusCode != 0 {
		return ctx.JSON(err.StatusCode, err)
	}
	return ctx.JSON(http.StatusCreated, response.Created(notification))
}

/*
 * @apiDefine: CreateNotificationResponse
 */
type CreateNotificationResponse struct {
	StatusCode int                 `json:"status_code" openapi:"example:201;type:integer;"`
	Data       models.Notification `json:"data" openapi:"$ref:Notification;type:object;"`
}
