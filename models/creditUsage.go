package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: CreditUsage
 */
type CreditUsage struct {
	ID                 uint                  `json:"id" gorm:"primaryKey;uniqueIndex:udx_credit_usages" openapi:"example:1"`
	PreviousCredit     int                   `json:"previous_credit" openapi:"example:1200"`
	CreditChangeAmount int                   `json:"credit_change_amount" openapi:"example:100"`
	ChangeType         string                `json:"change_type" openapi:"example:addition"` // addition or subtraction
	CurrentCredit      int                   `json:"current_credit" openapi:"example:1300"`
	StripeEventID      string                `json:"stripe_event_id" openapi:"example:evt_1J4J7v2eZvKYlo2Cz1z1"`
	UserID             uint                  `json:"user_id" openapi:"example:1"`
	User               *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" openapi:"$ref:UserData;type:object;"`
	CreatedAt          time.Time             `json:"created_at" openapi:"example:2021-01-01T00:00:00Z"`
	UpdatedAt          time.Time             `json:"updated_at" openapi:"example:2021-01-01T00:00:00Z"`
	DeletedAt          soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_credit_usages"`
}
