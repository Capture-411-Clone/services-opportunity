package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateNotificationRequestBody
 */
type UpdateNotificationRequestBody struct {
	Seen bool `json:"seen" openapi:"example:true"`
}

func (data *UpdateNotificationRequestBody) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"seen": gocriteria.New("seen").Optional(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"seen": "Seen",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, data)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

/*
 * @apiDefine: UpdateNotificationRequestParams
 */
type UpdateNotificationRequestParams struct {
	ID int `json:"id,string" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateNotificationRequestParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Update(ctx context.Context, id int, input UpdateNotificationRequestBody) (models.Notification, response.ErrorResponse) {

	var notification models.Notification

	err := s.db.WithContext(ctx).First(&notification, "id = ?", id).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Notification{}, response.GormErrorResponse(err, "Notification not found")
	}

	if !policy.CanDeleteNotification(ctx, notification) {
		s.logger.With(ctx).Error("You don't have permission to query notification")
		return models.Notification{}, response.ErrorBadRequest("You don't have permission to query notification")
	}

	err = s.db.WithContext(ctx).
		Model(&notification).
		Where("id = ?", id).
		Update("seen", input.Seen).
		Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return models.Notification{}, response.GormErrorResponse(nil, "An error occurred in update notification")
	}

	return notification, response.ErrorResponse{}
}
