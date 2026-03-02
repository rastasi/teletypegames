package domain

import (
	"gorm.io/gorm"
)

func MigrateGoDatabase(db *gorm.DB) {
	db.AutoMigrate(
		Software{},
		Release{},
	)
}
