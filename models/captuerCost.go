package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: CaptureCost
 */
type CaptureCost struct {
	ID                    uint                  `gorm:"primaryKey;uniqueIndex:udx_capture_costs" openapi:"example:1;type:integer;"`
	CompanyName           string                `json:"company_name" openapi:"example:Company Name;type:string;"`
	FullName              string                `json:"full_name" openapi:"example:James Mada;type:string;"`
	Email                 string                `json:"email" openapi:"example:4h7oK@example.com;type:string;"`
	ManagerCount          int                   `json:"manager_count" openapi:"example:5;type:integer;"`
	HourlyRate            int                   `json:"hourly_rate" openapi:"example:100;type:integer;"`
	AnnualPipelineTime    int                   `json:"annual_pipeline_time" openapi:"example:100;type:integer;"`
	CompanyAnnualCost     int                   `json:"company_annual_cost" openapi:"example:100;type:integer;"`
	CompanyCapture411Cost int                   `json:"company_capture411_cost" openapi:"example:100;type:integer;"`
	CompanySavings        int                   `json:"company_savings" openapi:"example:100;type:integer;"`
	CompanyHoursSaved     int                   `json:"company_hours_saved" openapi:"example:100;type:integer;"`
	CreatedAt             time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt             time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt             soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_capture_costs"`
}
