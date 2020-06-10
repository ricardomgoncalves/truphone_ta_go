package postgres

import (
	"github.com/jinzhu/gorm"
)

func CreateTables(db *gorm.DB) error {
	return db.AutoMigrate(&familyRow{}, &memberRow{}).Error
}