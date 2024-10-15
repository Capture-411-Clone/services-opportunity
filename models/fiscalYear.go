package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: FiscalYear
 */
type FiscalYear struct {
	ID        uint                  `gorm:"primaryKey;uniqueIndex:udx_fiscal_years;" openapi:"example:1;type:integer;"`
	User      *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID    int                   `json:"user_id" openapi:"example:1;type:integer;"`
	Year      string                `json:"year" gorm:"size:4;uniqueIndex:udx_fiscal_years;" openapi:"example:2022;type:string;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_fiscal_years;"`
}
