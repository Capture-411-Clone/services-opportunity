package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Market
 */
type Market struct {
	ID          uint                  `gorm:"primaryKey" openapi:"example:1;type:integer;"`
	Name        string                `json:"name" gorm:"size:200;uniqueIndex:udx_markets;" openapi:"example:Market Name;type:string;"`
	Acronym     string                `json:"acronym" gorm:"size:10;uniqueIndex:udx_markets;" openapi:"example:MN;type:string;"`
	User        *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID      int                   `json:"user_id" openapi:"example:1;type:integer;"`
	Departments []*Department         `json:"departments" openapi:"$ref:Department;type:array;"`
	CreatedAt   time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt   time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_markets"`
}
