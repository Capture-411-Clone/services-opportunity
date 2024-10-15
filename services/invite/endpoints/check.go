package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
)

/*
 * @apiDefine: CheckInviteRequestParams
 */
type CheckInviteRequestParams struct {
	Code string `json:"code" openapi:"example:invite_code;type:string"`
}

func (p *CheckInviteRequestParams) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"code": gocriteria.New("code").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"code": "Code",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Check(ctx context.Context, code string) (models.Invite, response.ErrorResponse) {
	var invite models.Invite

	err := s.db.WithContext(ctx).First(&invite, "code", code).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return invite, response.GormErrorResponse(nil, "invitation code not found")
	}

	return invite, response.ErrorResponse{}
}
