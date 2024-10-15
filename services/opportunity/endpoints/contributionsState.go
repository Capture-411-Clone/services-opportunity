package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/config"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

func (s *service) ContributionsState(ctx context.Context) (ContributionsStateResponseData, response.ErrorResponse) {

	var allCount int64
	err := s.db.Model(&models.Opportunity{}).Count(&allCount).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return ContributionsStateResponseData{}, response.ErrorInternalServerError(err, "cant process the request. contact support")
	}

	// var progressPercent float64
	// if allCount > 0 {
	// 	progressPercent = math.Floor(float64(allCount) / float64(config.AppConfig.ContributionThreshold) * 100)
	// }

	return ContributionsStateResponseData{
		Threshold: config.AppConfig.ContributionThreshold,
		// ProgressPercent: progressPercent,
		ProgressPercent: 2,
	}, response.ErrorResponse{}

}

/*
 * @apiDefine: ContributionsStateResponseData
 */
type ContributionsStateResponseData struct {
	Threshold       int     `json:"threshold"`
	ProgressPercent float64 `json:"progressPercent"`
}
