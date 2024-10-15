package seeders

import (
	"log"

	"github.com/Capture-411/services-opportunity/models"
	"github.com/Capture-411/services-opportunity/services/role"
	"gorm.io/gorm"
)

func RoleSeeder(db *gorm.DB) {

	var roles []models.Role
	for _, v := range role.Roles {
		roles = append(
			roles, models.Role{
				Title: v,
			},
		)
	}
	err := db.Create(&roles).Error
	if err != nil {
		log.Fatal(err)
	}
}
