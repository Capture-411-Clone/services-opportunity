package models

import (
	"time"

	"github.com/Capture-411/services-opportunity/kit/dtp"
	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Document
 */
type Document struct {
	ID        uint                  `gorm:"primaryKey;uniqueIndex:udx_documents" openapi:"example:1;type:integer;"`
	Title     string                `json:"title" gorm:"size:200;" openapi:"example:Document Title;type:string;"`
	FilePath  string                `json:"file_path" gorm:"size:200;" openapi:"example:Document File Path;type:string;"`
	Priority  int                   `json:"priority" openapi:"example:1;type:int;"`
	LinkType  string                `json:"link_type" openapi:"example:public"` // public , private
	UserID    int                   `json:"user_id" openapi:"example:1;type:int;"`
	User      User                  `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StaffID   dtp.NullInt64         `json:"staff_id" openapi:"example:1;type:int;"`
	Staff     *User                 `json:"staff" gorm:"foreignKey:StaffID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // staff
	OwnerID   int                   `json:"owner_id" openapi:"example:1;type:int;"`
	OwnerType string                `json:"owner_type" openapi:"example:Opportunity;type:string;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_documents"`
}
