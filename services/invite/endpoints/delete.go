package endpoints

import (
	"context"
	"net/http"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: DeleteInviteRequest
 */
type DeleteInviteRequest struct {
	IDs []int `json:"ids" openapi:"example:/[/1/,/2/,/3/];type:array;"`
}

func (a *DeleteInviteRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"ids": gocriteria.New("ids").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"ids": "IDs",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, a)
	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}
func (s *service) Delete(ctx context.Context, input DeleteInviteRequest) ([]int, response.ErrorResponse) {
	var res []int
	if !policy.CanDeleteInvite(ctx) {
		s.logger.With(ctx).Error("you don't have permission to delete invite")
		return res, response.ErrorForbidden("you don't have permission to delete invite")
	}

	err := s.db.WithContext(ctx).
		Where("id IN ?", input.IDs).
		Delete(&models.Invite{}).Error
	if err != nil {
		return res, response.GormErrorResponse(err, "an error occurred while deleting the invite")
	}
	return input.IDs, response.ErrorResponse{}
}
