package endpoints

import (
	"context"
	"strconv"
	"time"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	emailtemplates "github.com/Capture-411/services-opportunity/notification-v2/email/email-templates"
	"github.com/stripe/stripe-go/v76"
	"gorm.io/gorm"
)

func (s *service) HandleAddMoreCredits(ctx context.Context, cust *stripe.Customer, quantity int64, amountPaid int64, eventID string) response.ErrorResponse {
	// find credit usage with event id if it exists then return
	var creditUsage models.CreditUsage
	errr := s.db.Model(&models.CreditUsage{}).Where("stripe_event_id = ?", eventID).First(&creditUsage).Error
	if errr == nil {
		return response.ErrorResponse{}
	}

	var user models.User
	err := s.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&user, "stripe_customer_id = ?", cust.ID).Error
		if err != nil {
			s.logger.Errorf("Error finding user: %v\n", err)
			return err
		}

		creditUsage := models.CreditUsage{
			UserID:             user.ID,
			PreviousCredit:     user.Credits,
			ChangeType:         "addition",
			StripeEventID:      eventID,
			CreditChangeAmount: int(quantity),
			CurrentCredit:      user.Credits + int(quantity),
		}

		user.Credits = user.Credits + int(quantity)

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

		return nil
	})

	if err != nil {
		s.logger.Errorf("Error in transaction: %v\n", err)
		return response.ErrorBadRequest(nil, "Error in transaction")
	}

	quantityStr := strconv.Itoa(int(quantity))
	purchaseDateTodya := time.Now().Format("2006-01-02")
	amountPaidStr := strconv.Itoa(int(amountPaid))

	go s.sendCreditsAddedNotification(ctx, user, emailtemplates.CreditsAddedData{
		UserFullName:     user.Name + " " + user.LastName,
		CreditsPurchased: quantityStr,
		AmountPaid:       amountPaidStr,
		PurchaseDate:     purchaseDateTodya,
	})

	return response.ErrorResponse{}
}

func (s *service) sendCreditsAddedNotification(ctx context.Context, user models.User, data emailtemplates.CreditsAddedData) {
	var targetUsers []models.User
	targetUsers = append(targetUsers, user)

	subject, htmlMessage, err := emailtemplates.CreditsAddedTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	subjectShort, messageShort, err := emailtemplates.CreditsAddedShortTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	// Send notification for user
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"email"}, "", htmlMessage, subject)
	s.notifier.SendNotification(ctx, targetUsers, nil, []string{"database"}, messageShort, "", subjectShort)
}
