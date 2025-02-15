package util

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB, models ...any) {
	for _, model := range models {
		db := db.AutoMigrate(model)
		if db.Error != nil {
			panic(db.Error)
		}
	}
}
