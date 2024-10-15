package models

import (
	"time"

	"github.com/Capture-411/services-opportunity/kit/dtp"
	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Order
 */
type Order struct {
	ID            uint                  `gorm:"primaryKey;uniqueIndex:udx_agencies" openapi:"example:1;type:integer;"`
	FailedReason  string                `json:"failed_reason" gorm:"size:200" openapi:"example:Failed Reason;type:string;"`
	PriceAmount   float64               `json:"price_amount" openapi:"example:100.00;type:number;"`
	PaidAt        dtp.NullTime          `json:"paid_at" openapi:"example:2022-01-01T00:00:00Z"`
	OpportunityID int64                 `json:"opportunity_id" openapi:"example:1;type:integer;"`
	Opportunity   *Opportunity          `json:"opportunity" gorm:"foreignKey:OpportunityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:Opportunity;type:object;"`
	StripeEventID string                `json:"stripe_event_id" openapi:"example:evt_1J4J7v2eZvKYlo2Cz1z1"`
	RefundedAt    dtp.NullTime          `json:"refunded_at" openapi:"example:2022-01-01T00:00:00Z"`
	User          *User                 `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" openapi:"$ref:User;type:object;"`
	UserID        int                   `json:"user_id" openapi:"example:1;type:integer;"`
	CreatedAt     time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt     time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt     soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_agencies"`
}
