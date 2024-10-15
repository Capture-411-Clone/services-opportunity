package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

/*
 * @apiDefine: PassiveRevenue
 */
type PassiveRevenue struct {
	ID              uint                  `gorm:"primaryKey;uniqueIndex:udx_passive_revenues" openapi:"example:1;type:integer;"`
	CompanyName     string                `json:"company_name" openapi:"example:Company Name;type:string;"`
	FullName        string                `json:"full_name" openapi:"example:James Mada;type:string;"`
	Email           string                `json:"email" openapi:"example:4h7oK@example.com;type:string;"`
	LevelOneCount   int                   `json:"level_one_count" openapi:"example:5;type:integer;"`
	LevelTwoCount   int                   `json:"level_two_count" openapi:"example:5;type:integer;"`
	LevelThreeCount int                   `json:"level_three_count" openapi:"example:5;type:integer;"`
	LevelFourCount  int                   `json:"level_four_count" openapi:"example:5;type:integer;"`
	LevelFiveCount  int                   `json:"level_five_count" openapi:"example:5;type:integer;"`
	CreatedAt       time.Time             `json:"created_at" openapi:"example:2022-01-01T00:00:00Z"`
	UpdatedAt       time.Time             `json:"updated_at" openapi:"example:2022-01-01T00:00:00Z"`
	DeletedAt       soft_delete.DeletedAt `json:"deleted_at" gorm:"uniqueIndex:udx_passive_revenues"`
}
