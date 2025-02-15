package psql

import "gorm.io/gorm"

// MigrateModels can be used to run AutoMigrate on all models
func migrateModels(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}
