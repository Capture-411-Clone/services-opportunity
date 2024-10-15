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
 * @apiDefine: CreateNotificationRequest
 */
type CreateNotificationRequest struct {
	Body          string `json:"body" openapi:"example:Hi there from Capture411;"`
	Driver        string `json:"driver" openapi:"example:email"`
	TargetUserIDs string `json:"target_user_ids" openapi:"example:1"`
	Subject       string `json:"subject" openapi:"example:Hi there from Capture411;"`
}

func (data *CreateNotificationRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"body":            gocriteria.New("body").Required().MaxLength(500),
		"driver":          gocriteria.New("driver").Required(),
		"target_user_ids": gocriteria.New("target_user_ids").Required(),
		"subject":         gocriteria.New("subject").Required().MaxLength(500),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"body":            "Body",
			"driver":          "Driver",
			"target_user_ids": "Target User IDs",
			"subject":         "Subject",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, data)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) Create(ctx context.Context, input CreateNotificationRequest) ([]models.Notification, response.ErrorResponse) {

	if !policy.CanCreateNotification(ctx) {
		s.logger.With(ctx).Error("You don't have permission to create notification")
		return []models.Notification{}, response.ErrorForbidden("You don't have permission to create notification")
	}

	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var targetUsers []models.User

	drivers := strings.Split(input.Driver, ",")

	userIDsString := strings.Split(input.TargetUserIDs, ",")
	var targetUserIDs []int
	for _, idStr := range userIDsString {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return []models.Notification{}, response.GormErrorResponse(err, "An error occurred in find user")
		}
		targetUserIDs = append(targetUserIDs, id)
	}

	err := s.db.WithContext(ctx).Where("id IN ?", targetUserIDs).Find(&targetUsers).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Notification{}, response.GormErrorResponse(err, "An error occurred in find user")
	}

	var senderUser models.User
	err = s.db.WithContext(ctx).Where("id = ?", theID).First(&senderUser).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Notification{}, response.GormErrorResponse(err, "An error occurred in find user")
	}
	s.notifier.SendNotification(ctx, targetUsers, &senderUser, drivers, input.Body, input.Body, input.Subject)

	var notifications []models.Notification

	err = s.db.WithContext(ctx).Where("target_user_id IN ?", targetUserIDs).Find(&notifications).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return []models.Notification{}, response.GormErrorResponse(err, "An error occurred in find notification")
	}

	return notifications, response.ErrorResponse{}
}
