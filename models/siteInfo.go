package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: SiteInfo
 */
type SiteInfo struct {
	ID               uint                  `gorm:"primaryKey;uniqueIndex:udx_site_infos" openapi:"example:1;type:integer;"`
	CompanyName      string                `json:"company_name" openapi:"example:Company Name;type:string;"`
	MissionStatement string                `json:"mission_statement" openapi:"example:Mission Statement;type:string;"`
	History          string                `json:"history" openapi:"example:History;type:string;"`
	Goal             string                `json:"goal" openapi:"example:Goal;type:string;"`
	Value            string                `json:"value" openapi:"example:Value;type:string;"`
	Achievement      string                `json:"achievement" openapi:"example:Achievement;type:string;"`
	Member           string                `json:"member" openapi:"example:Member;type:string;"`
	GeneralContact   string                `json:"general_contact" openapi:"example:General Contact;type:string;"`
	Address          string                `json:"address" openapi:"example:Address;type:string;"`
	SocialMedia      string                `json:"social_media" openapi:"example:Social Media;type:string;"`
	PhoneNumber      string                `json:"phone_number" openapi:"example:Phone Number;type:string;"`
	EmailAddress     string                `json:"email_address" openapi:"example:Email Address;type:string;"`
	OfficeHours      string                `json:"office_hours" openapi:"example:Office Hours;type:string;"`
	PhysicalAddress  string                `json:"physical_address" openapi:"example:Physical Address;type:string;"`
	MapOrDirections  string                `json:"map_or_directions" openapi:"example:Map or Directions;type:string;"`
	EmergencyContact string                `json:"emergency_contact" openapi:"example:Emergency Contact;type:string;"`
	FeedbackLink     string                `json:"feedback_link" openapi:"example:Feedback Link;type:string;"`
	SupportLink      string                `json:"support_link" openapi:"example:Support Link;type:string;"`
	CreatedAt        time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt        time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt        soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_site_infos"`
}
