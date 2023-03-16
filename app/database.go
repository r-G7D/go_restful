package app

import (
	"r-G7D/go_restful/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DefDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("go_restful.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//TODO: db polling
	db.AutoMigrate(&domain.User{})

	return db
}
