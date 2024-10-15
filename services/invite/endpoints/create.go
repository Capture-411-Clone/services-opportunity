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
 * @apiDefine: CreateInviteRequest
 */
type CreateInviteRequest struct {
	Code   string `json:"code,omitempty" openapi:"example:invite_code;type:string"`
	Limit  int    `json:"limit,omitempty" openapi:"example:10;type:integer"`
	UserID int    `json:"user_id,omitempty" openapi:"example:1;type:integer"`
}

func (r *CreateInviteRequest) Validate(req *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"code":    gocriteria.New("code").Required(),
		"limit":   gocriteria.New("limit").Required(),
		"user_id": gocriteria.New("user_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"code":    "Code",
			"limit":   "Limit",
			"user_id": "User ID",
		},
	)

	errr := gocriteria.ValidateBody(req, schema, r)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Create(ctx context.Context, input CreateInviteRequest) (models.Invite, response.ErrorResponse) {
	var invite models.Invite
	if !policy.CanCreateInvite(ctx) {
		s.logger.With(ctx).Error("you don't have permission to create invite")
		return invite, response.ErrorForbidden("you don't have permission to create invite")
	}

	invite = models.Invite{
		Code:   dtp.NullString{String: input.Code, Valid: input.Code != ""},
		Limit:  input.Limit,
		UserID: input.UserID,
	}

	err := s.db.WithContext(ctx).Create(&invite).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return invite, response.GormErrorResponse(err, "an error occurred while creating the invite")
	}
	return invite, response.ErrorResponse{}
}
