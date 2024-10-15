package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Notification
 */
type Notification struct {
	ID           uint                  `json:"id" gorm:"primaryKey;uniqueIndex:udx_notifications" openapi:"example:1"`
	Body         string                `json:"body" gorm:"size:2000" openapi:"example:Hi there from Capture411"`
	Subject      string                `json:"subject" gorm:"size:2000" openapi:"example:Hi there from Capture411"`
	Recipient    string                `json:"recipient" gorm:"size:1000" openapi:"example:u9Ijz@example.com"`
	Driver       string                `json:"driver" openapi:"example:email"`
	Seen         bool                  `json:"seen" openapi:"example:true"`
	SenderUserID *uint                 `json:"sender_user_id" openapi:"example:1"`
	SenderUser   *User                 `json:"sender_user" gorm:"foreignKey:SenderUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" openapi:"$ref:UserData;type:object;"`
	TargetUserID *uint                 `json:"target_user_id" openapi:"example:1"`
	TargetUser   *User                 `json:"target_user" gorm:"foreignKey:TargetUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" openapi:"$ref:UserData;type:object;"`
	CreatedAt    time.Time             `json:"created_at" openapi:"example:2021-01-01T00:00:00Z"`
	UpdatedAt    time.Time             `json:"updated_at" openapi:"example:2021-01-01T00:00:00Z"`
	DeletedAt    soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_notifications"`
}
