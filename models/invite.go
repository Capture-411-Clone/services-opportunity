package models

import (
	"time"

	"github.com/Capture-411/services-opportunity/kit/dtp"
	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Invite
 */
type Invite struct {
	ID        uint                  `gorm:"primaryKey;" openapi:"example:1;type:integer;"`
	Code      dtp.NullString        `json:"code" gorm:"uniqueIndex:udx_invites" openapi:"example:invite_code;type:string;"`
	Limit     int                   `json:"limit" openapi:"example:200;type:int;"`
	UserID    int                   `json:"user_id" openapi:"example:1;type:int;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_invites"`
}
