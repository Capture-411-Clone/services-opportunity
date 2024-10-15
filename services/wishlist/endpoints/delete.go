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
 * @apiDefine: DeleteWishlistRequest
 */
type DeleteWishlistRequest struct {
	IDs string `json:"ids" openapi:"example:1;type:array;"`
}

func (a *DeleteWishlistRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
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

func (s *service) Delete(ctx context.Context, input DeleteWishlistRequest) (string, response.ErrorResponse) {
	Id := policy.ExtractIdClaim(ctx)
	userID, err := strconv.Atoi(Id)
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.ErrorBadRequest(nil, "Something went wrong")
	}

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
	var wishes []models.Wishlist
	err = s.db.WithContext(ctx).Where("id IN ? AND user_id = ?", ids, userID).Find(&wishes).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete from wishlist")
	}
	if len(wishes) == 0 {
		return "", response.ErrorResponse{
			Message: "Not found any from wishlist",
		}
	}

	err = s.db.WithContext(ctx).Where("id IN ?", ids).Delete(&models.Wishlist{}).Error

	if err != nil {
		s.logger.With(ctx).Error(err)
		return "", response.GormErrorResponse(err, "An error occurred in delete from wishlist")
	}

	return input.IDs, response.ErrorResponse{}
}
