package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Role
 */
type Role struct {
	ID        uint                  `gorm:"primaryKey;uniqueIndex:udx_roles" openapi:"example:1;type:integer;"`
	Title     string                `json:"title" gorm:"size:100;uniqueIndex:udx_role_name" openapi:"example:Role Title;type:string;"`
	Users     []*User               `json:"users" gorm:"many2many:user_role;" openapi:"$ref:User;type:array;"`
	CreatedAt time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_roles"`
}
