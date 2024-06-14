package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConectaDB(models ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	for _, model := range models {
		db.AutoMigrate(model)
	}

	return db, nil
}
