package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) Count(ctx context.Context, params EntityHuntQueryRequestParams) (int64, response.ErrorResponse) {
	var count int64

	tx := s.db.WithContext(ctx).Model(&models.EntityHunt{})

	if params.LastVisitedID != "" {
		tx = tx.Where("id > ?", params.LastVisitedID)
	}

	if params.SolicitationNumber != "" {
		tx = tx.Where("solicitation_number = ?", params.SolicitationNumber)
	}

	if params.Year != "" {
		tx = tx.Where("year = ?", params.Year)
	}

	if params.FileName != "" {
		tx = tx.Where("file_name = ?", params.FileName)
	}

	if params.FilePath != "" {
		tx = tx.Where("file_path = ?", params.FilePath)
	}

	err := tx.Count(&count).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return count, response.GormErrorResponse(err, "error in calculate data count")
	}
	return count, response.ErrorResponse{}
}
