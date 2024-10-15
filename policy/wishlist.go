package policy

import (
	"context"
)

func CanAccessWishlist(ctx context.Context) bool {
	return true
}
