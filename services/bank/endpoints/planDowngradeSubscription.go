package endpoints

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	emailtemplates "github.com/Capture-411/services-opportunity/notification-v2/email/email-templates"
	"github.com/Capture-411/services-opportunity/policy"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/subscription"
)

type PlanDowngradeRequestBody struct {
	NewPriceID string `json:"new_price_id"`
}

func (data *PlanDowngradeRequestBody) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"new_price_id": gocriteria.New("new_price_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"new_price_id": "New Price ID",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, data)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) PlanDowngradeSubscription(ctx context.Context, input PlanDowngradeRequestBody) response.ErrorResponse {
	id := policy.ExtractIdClaim(ctx)
	theID, _ := strconv.Atoi(id)

	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", theID).First(&user).Error
	if err != nil {
		s.logger.With(ctx).Errorf("CreateCreditCheckout: %v", err)
		return response.ErrorBadRequest(nil, "Could not proceed downgrading subscription")
	}

	if user.SubscriptionID == "" {
		return response.ErrorBadRequest("User does not have a subscription")
	}

	// Retrieve the current subscription to get its items, assuming it has only one item.
	currentSub, err := subscription.Get(user.SubscriptionID, nil)
	if err != nil {
		s.logger.Errorf("Error retrieving subscription: %v\n", err)
		return response.ErrorBadRequest("Error retrieving subscription")
	}

	// Check if the subscription is active or in a state that can be updated
	if currentSub.Status != stripe.SubscriptionStatusActive {
		// Handle cases where the subscription is not active or in an updatable state
		// This might involve reactivating the subscription or handling the error differently
		s.logger.Errorf("Subscription is not in an updatable state: Status=%v\n", currentSub.Status)
		return response.ErrorBadRequest("Subscription cannot be updated in its current state")
	}

	currentPrice := currentSub.Items.Data[0].Price
	currentPriceAmount := currentPrice.UnitAmount

	// get new price unit amount
	newPrice, err := price.Get(input.NewPriceID, nil)
	if err != nil {
		s.logger.With(ctx).Errorf("price.Get: %v", err)
		return response.ErrorBadRequest(nil, "Could not proceed with payment proccess")
	}

	newPriceAmount := newPrice.UnitAmount

	if newPriceAmount > currentPriceAmount {
		s.logger.Errorf("New price is greater than the current price: NewPrice=%v, CurrentPrice=%v\n", newPriceAmount, currentPriceAmount)
		return response.ErrorBadRequest("New price is greater than the current price")
	}

	// prevent user from downgrading twice
	err = s.db.Model(&models.Downgrade{}).Where("user_id = ? AND new_price_id = ?", theID, input.NewPriceID).First(&models.Downgrade{}).Error
	if err == nil {
		return response.ErrorBadRequest(nil, "You have already downgraded to this plan")
	}

	downgrade := models.Downgrade{
		UserID:         uint(theID),
		SubscriptionID: user.SubscriptionID,
		NewPriceID:     input.NewPriceID,
		NewProductID:   newPrice.Product.ID,
		DowngradeAt:    currentSub.CurrentPeriodEnd,
		NewProductName: newPrice.Product.Name,
	}

	err = s.db.Create(&downgrade).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error creating downgrade: %v", err)
		return response.ErrorBadRequest(nil, "Could not proceed downgrading subscription")
	}

	downgradeEffectiveDate := time.Unix(currentSub.CurrentPeriodEnd, 0)
	downgradeEffectiveDateString := downgradeEffectiveDate.Format("2006-01-02")

	go s.sendPlanDowngradeNotification(ctx, user, emailtemplates.SubscriptionDowngradedData{
		UserFullName:           user.Name + " " + user.LastName,
		PrevSubscriptionType:   currentPrice.Metadata["package_name"],
		NewSubscrptionType:     newPrice.Metadata["package_name"],
		DowngradeEffectiveDate: downgradeEffectiveDateString,
	})

	return response.ErrorResponse{}
}

func (s *service) sendPlanDowngradeNotification(ctx context.Context, user models.User, data emailtemplates.SubscriptionDowngradedData) {
	var targetUsers []models.User
	targetUsers = append(targetUsers, user)

	subject, htmlMessage, err := emailtemplates.SubscriptionDowngradedTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	subjectShort, messageShort, err := emailtemplates.SubscriptionDowngradedShortTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	// Send notification for user
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"email"}, "", htmlMessage, subject)
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"database"}, messageShort, "", subjectShort)
}
