package endpoints

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	emailtemplates "github.com/Capture-411/services-opportunity/notification-v2/email/email-templates"
	"github.com/stripe/stripe-go/v76"
)

func (s *service) HandleSubscriptionUpdated(ctx context.Context, cust *stripe.Customer, sub *stripe.Subscription) response.ErrorResponse {
	var user models.User
	err := s.db.First(&user, "stripe_customer_id = ?", cust.ID).Error
	if err != nil {
		s.logger.Errorf("Error finding user: %v\n", err)
		return response.ErrorBadRequest(nil, "User not found")
	}

	// Serialize the stripe.Subscription to a JSON string
	subJSON, err := json.Marshal(sub)
	if err != nil {
		s.logger.Errorf("Error serializing subscription: %v\n", err)
		return response.ErrorBadRequest(nil, "Error serializing subscription")
	}

	user.Subscription = subJSON
	user.SubscriptionID = sub.ID

	err = s.db.Save(&user).Error
	if err != nil {
		s.logger.Errorf("Error saving user: %v\n", err)
		return response.ErrorBadRequest(nil, "Error saving user")
	}

	cancelAtDate := time.Unix(sub.CancelAt, 0)
	cancelAtString := cancelAtDate.Format("2006-01-02")

	if sub.CancelAtPeriodEnd && sub.CanceledAt != 0 {
		go s.sendHandleSubscriptionCanceledNotification(ctx, user, emailtemplates.SubscriptionCanceledData{
			UserFullName:  user.Name + " " + user.LastName,
			EffectiveDate: cancelAtString,
		})
	}

	return response.ErrorResponse{}
}

func (s *service) sendHandleSubscriptionCanceledNotification(ctx context.Context, user models.User, data emailtemplates.SubscriptionCanceledData) {
	var targetUsers []models.User
	targetUsers = append(targetUsers, user)

	subject, htmlMessage, err := emailtemplates.SubscriptionCanceledTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	subjectShort, messageShort, err := emailtemplates.SubscriptionCanceledShortTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	// Send notification for user
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"email"}, "", htmlMessage, subject)
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"database"}, messageShort, "", subjectShort)
}
