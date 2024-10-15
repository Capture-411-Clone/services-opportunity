package endpoints

import (
	"context"
	"errors"

	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

func (s *service) findUser(ctx context.Context, username string) (bool, models.User, error) {
	var count int64
	var exists bool
	var user models.User
	err := s.db.WithContext(ctx).
		Where("LOWER(phone) = LOWER(?)", username).
		Or("LOWER(email) = LOWER(?)", username).
		Or("LOWER(username) = LOWER(?)", username).
		Preload("Roles").
		Preload("Invite").
		First(&user).Count(&count).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, user, nil
	}
	if err != nil {
		return false, user, err
	}
	exists = count > 0
	return exists, user, nil
}
