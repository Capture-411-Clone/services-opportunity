package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Department
 */
type Department struct {
	ID        uint                  `gorm:"primaryKey;" openapi:"example:1;type:integer;"`
	Name      string                `json:"name" gorm:"size:200;uniqueIndex:udx_departments;" openapi:"example:Department Name;type:string;"`
	Acronym   string                `json:"acronym" gorm:"size:10;" openapi:"example:DEP;type:string;"`
	Market    *Market               `json:"market" gorm:"foreignKey:MarketID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:Market;type:object;"`
	MarketID  int                   `json:"market_id" openapi:"example:1;type:integer;"`
	Agencies  []*Agency             `json:"agencies" openapi:"$ref:Agency;type:array;"`
	User      *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID    int                   `json:"user_id" openapi:"example:1;type:integer;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_departments"`
}
