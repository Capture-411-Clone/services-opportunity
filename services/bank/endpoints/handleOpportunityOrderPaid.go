package endpoints

import (
	"context"
	"fmt"
	"time"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	emailtemplates "github.com/Capture-411/services-opportunity/notification-v2/email/email-templates"
	"github.com/stripe/stripe-go/v76"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *service) HandleOpportunityOrderPaid(ctx context.Context, cust *stripe.Customer, orderID int64, eventID string) response.ErrorResponse {

	// find order with event id if it exists then return
	var order models.Order
	errr := s.db.Model(&models.Order{}).Where("stripe_event_id = ?", eventID).First(&order).Error
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

		// Get the opportunity order
		var order models.Order
		err = tx.Model(&models.Order{}).Where("id = ?", orderID).Update("paid_at", time.Now()).Preload("Opportunity").First(&order).Error
		if err != nil {
			s.logger.Errorf("Error finding or updating order: %v\n", err)
			return err
		}

		order.StripeEventID = eventID
		err = tx.Save(&order).Error
		if err != nil {
			s.logger.Errorf("Error saving order: %v\n", err)
			return err
		}

		// remove from item from that user's whishlist
		err = tx.Model(&models.Wishlist{}).Where("user_id = ? AND opportunity_id = ?", user.ID, order.OpportunityID).Delete(&models.Wishlist{}).Error
		if err != nil {
			s.logger.Errorf("Error removing item from wishlist: %v\n", err)
			return err
		}

		amountPaidStr := fmt.Sprintf("%.2f", float64(order.PriceAmount)/100)
		opportunityOrderedDateString := order.CreatedAt.Format("2006-01-02")
		s.sendOpportunityOrderActivatedNotification(ctx, user, emailtemplates.OpportunityOrderPaidData{
			UserFullName:                  user.Name + " " + user.LastName,
			AmountPaid:                    amountPaidStr,
			OpportunityOrderedDate:        opportunityOrderedDateString,
			OpportunityTitle:              order.Opportunity.Title,
			OpportunitySolicitationNumber: order.Opportunity.SolicitationNumber,
		})

		return nil
	})

	if err != nil {
		s.logger.Errorf("Error in transaction: %v\n", err)
		return response.ErrorBadRequest(nil, "Could not proceed with payment process")
	}
	return response.ErrorResponse{}
}

func (s *service) sendOpportunityOrderActivatedNotification(ctx context.Context, user models.User, data emailtemplates.OpportunityOrderPaidData) {
	var targetUsers []models.User
	targetUsers = append(targetUsers, user)

	subject, plainMessage, htmlMessage, err := emailtemplates.OpportunityOrderPaidTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	subjectShort, messageShort, err := emailtemplates.OpportunityOrderPaidShortTemplate(data)
	if err != nil {
		s.logger.Errorf("Error getting email template: %v\n", err)
		return
	}

	err = s.notifier.SendNotification(ctx, targetUsers, nil, []string{"email"}, plainMessage, htmlMessage, subject)
	if err != nil {
		s.logger.Errorf("Error sending email notification: %v\n", err)
	}

	err = s.notifier.SendNotification(ctx, targetUsers, nil, []string{"database"}, messageShort, "", subjectShort)
	if err != nil {
		s.logger.Errorf("Error sending database notification: %v\n", err)
	}
}
