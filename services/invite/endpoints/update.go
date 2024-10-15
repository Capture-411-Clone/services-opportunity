package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/dtp"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: UpdateInviteRequest
 */
type UpdateInviteRequest struct {
	Code  string `json:"code,omitempty" openapi:"example:1;type:string"`
	Limit int    `json:"limit,omitempty" openapi:"example:10;type:integer"`
}

func (r *UpdateInviteRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"code":  gocriteria.New("code").Required(),
		"limit": gocriteria.New("limit").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"code":  "Code",
			"limit": "Limit",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

/*
 * @apiDefine: UpdateInviteRequestParam
 */
type UpdateInviteRequestParam struct {
	ID string `json:"id" openapi:"example:1;nullable;pattern:^[0-9]+$;in:path"`
}

func (p *UpdateInviteRequestParam) Validate(params gocriteria.Params) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"id": gocriteria.New("id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"id": "ID",
		},
	)

	errr := gocriteria.ValidateParams(params, schema, p)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Update(ctx context.Context, id string, input UpdateInviteRequest) (models.Invite, response.ErrorResponse) {
	var invite models.Invite
	if !policy.CanUpdateInvite(ctx) {
		s.logger.With(ctx).Error("you don't have permission to update invite")
		return invite, response.ErrorForbidden("you don't have permission to update invite")
	}
	err := s.db.WithContext(ctx).First(&invite, "id", id).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return invite, response.GormErrorResponse(err, "invite code not found")
	}
	invite.Limit = input.Limit
	invite.Code = dtp.NullString{String: input.Code, Valid: input.Code != ""}

	err = s.db.WithContext(ctx).Save(&invite).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return invite, response.GormErrorResponse(err, "error in updating the invite code")
	}

	return invite, response.ErrorResponse{}
}
