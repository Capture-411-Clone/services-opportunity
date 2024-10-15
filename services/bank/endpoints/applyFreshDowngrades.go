package endpoints

import (
	"context"
	"errors"
	"time"

	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) ApplyFreshDowngrades(ctx context.Context) error {

	var downgrades []models.Downgrade
	// downgrade_at timestamp less or equal to current time
	err := s.db.Model(&models.Downgrade{}).Where("downgrade_at <= ?", time.Now().Unix()).Find(&downgrades).Error
	if err != nil {
		s.logger.Errorf("Error fetching downgrades: %v", err)
		return errors.New("could not find downgrades")
	}

	for _, downgrade := range downgrades {
		downgradeData := ChangeSubscriptionForDowngradeRequestBody{
			DowngradeID: downgrade.ID,
			UserID:      downgrade.UserID,
		}

		s.logger.Infof("Downgrading id=[%d] name=[%s] plan...", downgrade.ID, downgrade.NewProductName)

		_, err := s.ChangeSubscriptionForDowngrade(ctx, downgradeData)
		if err.StatusCode != 0 {
			s.logger.With(ctx).Errorf("Error applying downgrade: %v", err)
		}
	}

	return nil
}
