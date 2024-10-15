package endpoints

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Capture-411/services-opportunity/gocriteria"
	"github.com/Capture-411/services-opportunity/gocriteria/message"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/policy"
)

/*
 * @apiDefine: AddWishlistRequest
 */
type AddWishlistRequest struct {
	OpportunityID int `json:"opportunity_id" openapi:"example:1;type:integer"`
}

func (c *AddWishlistRequest) Validate(r *http.Request) gocriteria.ValidityResponseErrors {
	schema := gocriteria.Schema{
		"opportunity_id": gocriteria.New("opportunity_id").Required(),
	}

	gocriteria.SetFieldLabels(
		&message.Labels{
			"opportunity_id": "Opportunity ID",
		},
	)

	errr := gocriteria.ValidateBody(r, schema, c)

	if len(errr) > 0 {
		dumpedErrors := gocriteria.DumpErrors(errr)
		return dumpedErrors
	}

	return nil
}

func (s *service) Add(ctx context.Context, input AddWishlistRequest) (models.Wishlist, response.ErrorResponse) {
	Id := policy.ExtractIdClaim(ctx)
	id, _ := strconv.Atoi(Id)

	wishlist := models.Wishlist{
		UserID:        id,
		OpportunityID: input.OpportunityID,
	}
	err := s.db.WithContext(ctx).Create(&wishlist).Error
	if err != nil {
		s.logger.With(ctx).Error(err)
		return wishlist, response.GormErrorResponse(err, "An error occurred while creating the wishlist")
	}
	return wishlist, response.ErrorResponse{}
}
