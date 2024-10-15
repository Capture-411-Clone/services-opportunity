package notificationv2

import (
	"context"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/notification-v2/email"
	"gorm.io/gorm"
)

type Notifier interface {
	SendNotification(ctx context.Context, targetUsers []models.User, sender *models.User, drivers []string, plainMessage, htmlMessage, subject string) error
}

type notifier struct {
	db *gorm.DB
}

func MakeNotifier(db *gorm.DB) Notifier {
	return &notifier{
		db: db,
	}
}

func (n notifier) SendNotification(ctx context.Context, targetUsers []models.User, sender *models.User, drivers []string, plainMessage string, htmlMessage string, subject string) error {
	if len(targetUsers) == 0 {
		return nil // No target users to process
	}

	var senderUserID *uint // A pointer to hold sender's user ID, can be nil
	if sender != nil {
		senderUserID = &sender.ID // Assign sender's ID if sender is not nil
	}

	for _, targetUser := range targetUsers {
		recipientEmail := targetUser.Email.String // Assuming targetUser.Email.String is valid

		var targetUserID *uint // A pointer to hold target user's ID
		if targetUser.ID != 0 {
			targetUserID = &targetUser.ID // Assign target user's ID if it's not zero
		}

		for _, driver := range drivers {
			notif := models.Notification{
				Body:         plainMessage,
				Subject:      subject,
				Recipient:    "",
				Driver:       driver,
				TargetUserID: targetUserID,
				SenderUserID: senderUserID,
			}

			// Handle email notifications
			if driver == "email" && config.AppConfig.SendEmail {
				if recipientEmail != "" {
					notif.Recipient = recipientEmail
					if err := email.Send(recipientEmail, subject, plainMessage, htmlMessage); err != nil {
						return err
					}
				}
			}

			// Save the notification in the database for each driver
			if err := n.db.WithContext(ctx).Create(&notif).Error; err != nil {
				return err
			}

		}
	}

	return nil
}
