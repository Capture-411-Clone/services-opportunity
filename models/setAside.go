package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: SetAside
 */
type SetAside struct {
	ID        uint                  `gorm:"primaryKey;uniqueIndex:udx_set_asides;" openapi:"example:1;type:integer;"`
	Name      string                `json:"name" gorm:"size:200;" openapi:"example:Set Aside Name;type:string;"`
	Acronym   string                `json:"acronym" gorm:"size:10;" openapi:"example:SA;type:string;"`
	User      *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID    int                   `json:"user_id" openapi:"example:1;type:integer;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_set_asides;"`
}
