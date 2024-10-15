package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: RegisterInvite
 */
type RegisterInvite struct {
	ID        uint                  `gorm:"primaryKey;uniqueIndex:udx_register_invites" openapi:"example:1;type:integer;"`
	HostID    int                   `json:"host_id" openapi:"example:1;type:integer;"`
	Host      User                  `json:"host" gorm:"foreignKey:HostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" openapi:"$ref:User;type:object;"`
	UserID    int                   `json:"user_id" openapi:"example:1;type:integer;"`
	User      User                  `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" openapi:"$ref:User;type:object;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_register_invites"`
}
