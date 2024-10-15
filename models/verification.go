package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: Verification
 */
type Verification struct {
	ID          uint                  `gorm:"primaryKey;uniqueIndex:udx_verifications"`
	Phone       string                `json:"phone" gorm:"size:20;index"`
	Email       string                `json:"email" gorm:"size:256;index"`
	Code        string                `json:"code" gorm:"size:10;"`
	SessionCode string                `json:"session_code" gorm:"size:256;index"`
	ExchangedAt time.Time             `json:"exchanged_at"`
	ExpiresAt   time.Time             `json:"expires_at"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_verifications"`
}

func (v *Verification) Expired() bool {
	return v.ExpiresAt.Before(time.Now())
}

func (v *Verification) Exchanged() bool {
	return !v.ExchangedAt.Before(time.Now())
}

func (v *Verification) NotValidPhone(phone string) bool {
	return v.Phone != phone
}

func (v *Verification) NotValidEmail(email string) bool {
	return v.Email != email
}
