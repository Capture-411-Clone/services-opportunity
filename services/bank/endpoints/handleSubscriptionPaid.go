package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	emailtemplates "github.com/Capture-411/services-opportunity/notification-v2/email/email-templates"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/subscription"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *service) HandleSubscriptionPaid(ctx context.Context, cust *stripe.Customer, sub *stripe.Subscription, price *stripe.Price, eventID string) response.ErrorResponse {
	// find credit usage with event id if it exists then return
	var creditUsage models.CreditUsage
	errr := s.db.Model(&models.CreditUsage{}).Where("stripe_event_id = ?", eventID).First(&creditUsage).Error
	if errr == nil {
		return response.ErrorResponse{}
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var user models.User

		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, "stripe_customer_id = ?", cust.ID).Error
		if err != nil {
			s.logger.Errorf("Error finding user: %v\n", err)
			return err
		}

		intCreditAmount, err := strconv.Atoi(price.Metadata["credit_amount"])
		if err != nil {
			s.logger.Errorf("Error converting credit amount to int: %v\n", err)
			return err
		}

		// Serialize the stripe.Subscription to a JSON string
		subJSON, err := json.Marshal(sub)
		if err != nil {
			s.logger.Errorf("Error serializing subscription: %v\n", err)
			return err
		}

		fmt.Println("********* Paind Subscription json data", string(subJSON))
		fmt.Printf("********* Price data %v", price)

		user.Subscription = subJSON
		user.SubscriptionID = sub.ID
		user.Credits = user.Credits + intCreditAmount

		// Prepare the creditUsage record
		creditUsage := models.CreditUsage{
			UserID:             user.ID,
			PreviousCredit:     user.Credits - intCreditAmount, // Assuming the credits before adding
			ChangeType:         "addition",
			StripeEventID:      eventID,
			CreditChangeAmount: intCreditAmount,
			CurrentCredit:      user.Credits,
		}

		err = tx.Save(&creditUsage).Error
		if err != nil {
			s.logger.Errorf("Error saving credit usage: %v\n", err)
			return err
		}

		err = tx.Save(&user).Error
		if err != nil {
			s.logger.Errorf("Error saving user: %v\n", err)
			return err
		}

		// Cancel all active subscriptions except for the last one paid
		if err := s.cancelActiveSubscriptionsExceptLast(cust.ID, sub.ID); err != nil {
			s.logger.Errorf("Error cancelling other active subscriptions: %v\n", err)
			return err
		}

		var downgrades []models.Downgrade
		err = tx.Model(&models.Downgrade{}).Where("user_id = ?", user.ID).Find(&downgrades).Error
		if err != nil {
			s.logger.With(ctx).Errorf("Error fetching downgrade: %v", err)
			return err
		}

		// Delete All downgrades
		err = tx.Model(&models.Downgrade{}).Where("user_id = ?", user.ID).Delete(&downgrades).Error
		if err != nil {
			s.logger.With(ctx).Errorf("Error deleting downgrade: %v", err)
			return err
		}

		subscriptionType := price.Metadata["package_name"]

		var recurring string
		interval := string(price.Recurring.Interval) // Convert enum to string

		switch interval {
		case "month":
			recurring = "Monthly"
		case "year":
			recurring = "Yearly"
		default:
			recurring = interval // Keeps the original value for any other intervals
		}

		amountPaid := fmt.Sprintf("$%.2f", float64(price.UnitAmount)/100.0)

		// Convert int64 Unix timestamps to time.Time
		start := time.Unix(sub.CurrentPeriodStart, 0)
		end := time.Unix(sub.CurrentPeriodEnd, 0)

		// Format time.Time objects to "YYYY-MM-DD"
		subscriptionDateString := fmt.Sprintf("%s to %s", start.Format("2006-01-02"), end.Format("2006-01-02"))

		go s.sendSubscriptionActivatedNotification(ctx, user, emailtemplates.SubscriptionPaidData{
			UserFullName:     user.Name + " " + user.LastName,
			SubscriptionType: subscriptionType + " " + recurring,
			CreditsAllocated: price.Metadata["credit_amount"],
			AmountPaid:       amountPaid,
			SubscriptionDate: subscriptionDateString,
		})

		return nil
	})

	if err != nil {
		s.logger.Errorf("Error in transaction: %v\n", err)
		return response.ErrorBadRequest(nil, "Could not handle subscription payment. Please contact support if this issue persists.")
	}
	return response.ErrorResponse{}
}

func (s *service) sendSubscriptionActivatedNotification(ctx context.Context, user models.User, data emailtemplates.SubscriptionPaidData) {
	var targetUsers []models.User
	targetUsers = append(targetUsers, user)

	subject, htmlMessage, err := emailtemplates.SubscriptionPaidTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	subjectShort, messageShort, err := emailtemplates.SubscriptionPaidShortTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"email"}, "", htmlMessage, subject)
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"database"}, messageShort, "", subjectShort)
}

// cancelActiveSubscriptionsExceptLast cancels all active subscriptions for the customer except for the most recently paid one.
func (s *service) cancelActiveSubscriptionsExceptLast(customerID, keepSubscriptionID string) error {

	params := &stripe.SubscriptionListParams{
		Customer: &customerID,
		Status:   stripe.String(string(stripe.SubscriptionStatusActive)),
	}
	params.AddExpand("data.default_payment_method")
	i := subscription.List(params)
	for i.Next() {
		activeSub := i.Subscription()
		// Skip the subscription that matches the keepSubscriptionID
		if activeSub.ID == keepSubscriptionID {
			continue
		}

		_, err := subscription.Cancel(activeSub.ID, nil)
		if err != nil {
			s.logger.Errorf("Failed to cancel subscription %v: %v", activeSub.ID, err)
			// Depending on your requirements, you might want to continue trying to cancel the other subscriptions
			// even if one cancellation fails, instead of returning immediately.
			return err
		}
	}

	if err := i.Err(); err != nil {
		return err
	}

	return nil
}
