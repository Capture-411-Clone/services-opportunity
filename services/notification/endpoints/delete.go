package endpoints

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: DeleteNotificationRequestBody
 */
type DeleteNotificationRequestBody struct {
	IDs string `json:"ids" openapi:"example:1,2,3;required;"`
}

func (a *DeleteNotificationRequestBody) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"ids": gocriteria.New("ids").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"ids": "IDs",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, a)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Delete(ctx context.Context, input DeleteNotificationRequestBody) (string, response.ErrorResponse) {
	var ids []int
	for _, idStr := range strings.Split(input.IDs, ",") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorResponse{
				Message: "Invalid input IDs",
			}
		}
		ids = append(ids, id)
	}

	// find first
	var notifications []models.Notification
	err := s.db.WithContext(ctx).Where("id IN ?", ids).Find(&notifications).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete notifications")
	}
	if len(notifications) == 0 {
		return "", response.ErrorResponse{
			Message: "Not found any notifications",
		}
	}

	for _, notification := range notifications {
		if !policy.CanDeleteNotification(ctx, notification) {
			s.logger.With(ctx).Error("You don't have permission to delete a notification")
			return "", response.ErrorBadRequest("You don't have permission to delete a notification")
		}
	}

	err = s.db.WithContext(ctx).Where("id IN ?", ids).Delete(&models.Notification{}).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete notifications")
	}

	return input.IDs, response.ErrorResponse{}
}
