package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Naics
 */
type Naics struct {
	ID        uint                  `gorm:"primaryKey;uniqueIndex:udx_naicses;" openapi:"example:1;type:integer;"`
	Name      string                `json:"name" gorm:"size:200;" openapi:"example:Naics Name;type:string;"`
	User      *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID    int                   `json:"user_id" openapi:"example:1;type:integer;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_naicses;"`
}
