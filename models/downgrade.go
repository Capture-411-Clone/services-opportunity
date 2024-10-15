package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Downgrade
 */
type Downgrade struct {
	ID             uint                  `json:"id" gorm:"primaryKey;uniqueIndex:udx_tags" openapi:"example:1"`
	UserID         uint                  `json:"user_id" openapi:"example:1"`
	User           *User                 `json:"user" openapi:"$ref:User;type:object;"`
	NewPriceID     string                `json:"new_price_id" openapi:"example:price_1J2j3K4L5M"`
	NewProductID   string                `json:"new_product_id" openapi:"example:prod_1A2B3C4D5E"`
	SubscriptionID string                `json:"subscription_id" openapi:"example:sub_1A2B3C4D5E"`
	NewProductName string                `json:"new_product_name" openapi:"example:product_name"`
	DowngradeAt    int64                 `json:"downgrade_at" openapi:"example:1612137600"`
	CreatedAt      time.Time             `json:"created_at" openapi:"example:2021-01-01T00:00:00Z"`
	UpdatedAt      time.Time             `json:"updated_at" openapi:"example:2021-01-01T00:00:00Z"`
	DeletedAt      soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_tags"`
}
