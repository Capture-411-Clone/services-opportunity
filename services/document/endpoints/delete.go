package endpoints

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: DeleteDocumentsRequest
 */
type DeleteDocumentsRequest struct {
	IDs string `json:"ids" openapi:"example:1,2,3;"`
}

func (a *DeleteDocumentsRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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
func (s *service) Delete(ctx context.Context, input DeleteDocumentsRequest) (string, response.ErrorResponse) {
	var ids []int
	for _, idStr := range strings.Split(input.IDs, ",") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.logger.With(ctx).Error(err)
			return "", response.ErrorResponse{
				Message: "Invalid input IDs",
			}
		}
		ids = append(ids, id)
	}

	// find first
	var documents []models.Document
	err := s.db.WithContext(ctx).Where("id IN ?", ids).Find(&documents).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete document")
	}
	if len(documents) == 0 {
		return "", response.ErrorResponse{
			Message: "Not found any document",
		}
	}

	if !policy.CanDeleteDocument(ctx, documents) {
		s.logger.With(ctx).Error("you do not have permission to remove documents a opportunity")
		return "", response.ErrorForbidden("you do not have permission to remove documents a opportunity")
	}

	err = s.db.WithContext(ctx).Where("id IN ?", ids).Delete(&models.Document{}).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete document")
	}

	return input.IDs, response.ErrorResponse{}
}
