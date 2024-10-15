package endpoints

import (
	"context"

	"github.com/Capture-411/services-opportunity/kit/log"
	"github.com/Capture-411/services-opportunity/kit/response"
	"github.com/Capture-411/services-opportunity/models"
	"gorm.io/gorm"
)

type Service interface {
	Query(ctx context.Context, offset, limit int, filters WishlistQueryRequestParams) (
		wishlists []models.Wishlist, err response.ErrorResponse,
	)
	Count(ctx context.Context, params WishlistQueryRequestParams) (count int64, err response.ErrorResponse)
	Add(ctx context.Context, input AddWishlistRequest) (wishlist models.Wishlist, err response.ErrorResponse)
	Delete(ctx context.Context, input DeleteWishlistRequest) (response string, err response.ErrorResponse)
}

type service struct {
	db     *gorm.DB
	logger log.Logger
}

func MakeService(db *gorm.DB, logger log.Logger) Service {
	return &service{
		db:     db,
		logger: logger,
	}
}
