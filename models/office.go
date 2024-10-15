package models

import (
	"time"

	"github.com/Capture-411/services-opportunity/kit/dtp"
	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Office
 */
type Office struct {
	ID        uint                  `gorm:"primaryKey;uniqueIndex:udx_offices;" openapi:"example:1;type:integer;"`
	Name      string                `json:"name" gorm:"size:200" openapi:"example:Office Name;type:string;"`
	Acronym   string                `json:"acronym" gorm:"size:10;" openapi:"example:OFF;type:string;"`
	Children  []*Office             `json:"children" gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:Office;type:array;"`
	Parent    *Office               `json:"parent" gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:Office;type:object;"`
	ParentID  dtp.NullInt64         `json:"parent_id" openapi:"example:1;type:integer;"`
	Agency    *Agency               `json:"agency" gorm:"foreignKey:AgencyID;constraint:OnUpdate:CASCADE;" openapi:"$ref:Agency;type:object;"`
	AgencyID  dtp.NullInt64         `json:"agency_id" openapi:"example:1;type:integer;"`
	User      *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID    int                   `json:"user_id" openapi:"example:1;type:integer;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_offices"`
}
