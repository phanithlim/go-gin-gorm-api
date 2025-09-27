package config

import (
	"go-gin-gorm-api/helper"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
