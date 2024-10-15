package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Wishlist
 */
type Wishlist struct {
	ID            uint                  `gorm:"primaryKey;uniqueIndex:udx_wishlists;" openapi:"example:1;type:integer;"`
	User          *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID        int                   `json:"user_id" openapi:"example:1;type:integer;"`
	Opportunity   *Opportunity          `json:"opportunity" gorm:"foreignKey:OpportunityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:Opportunity;type:object;"`
	OpportunityID int                   `json:"opportunity_id" openapi:"example:1;type:integer;"`
	CreatedAt     time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt     time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_wishlists;"`
}
