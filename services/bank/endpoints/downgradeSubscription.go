package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/subscription"
)

type ChangeSubscriptionForDowngradeRequestBody struct {
	DowngradeID uint `json:"downgrade_id"`
	UserID      uint `json:"user_id"`
}

func (data *ChangeSubscriptionForDowngradeRequestBody) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"downgrade_id": gocriteria.New("downgrade_id").Required(),
		"user_id":      gocriteria.New("user_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"downgrade_id": "Downgrade ID",
			"user_id":      "User ID",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, data)

	if len(errr) > 0 {
		return gocriteria.DumpErrors(errr)
	}

	return nil
}

func (s *service) ChangeSubscriptionForDowngrade(ctx context.Context, input ChangeSubscriptionForDowngradeRequestBody) (*stripe.Subscription, response.ErrorResponse) {
	var user models.User
	err := s.db.Model(&models.User{}).Where("id = ?", input.UserID).First(&user).Error
	if err != nil {
		s.logger.With(ctx).Errorf("CreateCreditCheckout: %v", err)
		return nil, response.ErrorBadRequest(nil, "Could not proceed downgrading subscription")
	}

	if user.SubscriptionID == "" {
		return nil, response.ErrorBadRequest("User does not have a subscription")
	}

	var downgrade models.Downgrade
	err = s.db.Model(&models.Downgrade{}).Where("id = ?", input.DowngradeID).First(&downgrade).Error
	if err != nil {
		s.logger.With(ctx).Errorf("CreateCreditCheckout: %v", err)
		return nil, response.ErrorBadRequest(nil, "Could not proceed downgrading subscription")
	}

	// Retrieve the current subscription to get its items, assuming it has only one item.
	currentSub, err := subscription.Get(downgrade.SubscriptionID, nil)
	if err != nil {
		s.logger.Errorf("Error retrieving subscription: %v\n", err)
		return nil, response.ErrorBadRequest("Error retrieving subscription")
	}

	// Check if the subscription is active or in a state that can be updated
	if currentSub.Status != stripe.SubscriptionStatusActive {
		// Handle cases where the subscription is not active or in an updatable state
		// This might involve reactivating the subscription or handling the error differently
		s.logger.Errorf("Subscription is not in an updatable state: Status=%v\n", currentSub.Status)
		return nil, response.ErrorBadRequest("Subscription cannot be updated in its current state")
	}

	currentPriceAmount := currentSub.Items.Data[0].Price.UnitAmount

	// get new price unit amount
	newPrice, err := price.Get(downgrade.NewPriceID, nil)
	if err != nil {
		s.logger.With(ctx).Errorf("price.Get: %v", err)
		return nil, response.ErrorBadRequest(nil, "Could not proceed with payment proccess")
	}

	newPriceAmount := newPrice.UnitAmount

	if newPriceAmount > currentPriceAmount {
		s.logger.Errorf("New price is greater than the current price: NewPrice=%v, CurrentPrice=%v\n", newPriceAmount, currentPriceAmount)
		return nil, response.ErrorBadRequest("New price is greater than the current price")
	}

	// Prepare the update parameters to change the price and prevent immediate proration.
	params := &stripe.SubscriptionParams{
		CancelAtPeriodEnd: stripe.Bool(false),
		ProrationBehavior: stripe.String(string(stripe.SubscriptionSchedulePhaseProrationBehaviorNone)),
		Items: []*stripe.SubscriptionItemsParams{
			{
				ID:    &currentSub.Items.Data[0].ID,
				Price: stripe.String(downgrade.NewPriceID), // Set the new lower price ID for the subscription item.
			},
		},
		AutomaticTax: &stripe.SubscriptionAutomaticTaxParams{
			Enabled: stripe.Bool(true),
		},
	}

	// Schedule the subscription update.
	updatedSub, err := subscription.Update(downgrade.SubscriptionID, params)
	if err != nil {
		s.logger.Errorf("Error updating subscription: %v\n", err)
		return nil, response.ErrorBadRequest("Error updating subscription")
	}

	// Remove From Downgrade Database
	err = s.db.Delete(&downgrade).Error
	if err != nil {
		s.logger.With(ctx).Errorf("Error deleting downgrade: %v", err)
		return nil, response.ErrorBadRequest(nil, "Could not proceed downgrading subscription")
	}

	go s.sendChangeSubscriptionForDowngradeNotification(ctx, user)

	return updatedSub, response.ErrorResponse{}
}

func (s *service) sendChangeSubscriptionForDowngradeNotification(ctx context.Context, user models.User) {
	var targetUsers []models.User
	targetUsers = append(targetUsers, user)
	// Send notification for user
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"email"}, "Your subscription has been downgraded", "", "Subscription Downgraded")
}
