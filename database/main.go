package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/anon767/swaggertest/domain"
)

func Bootstrap() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&domain.Person{})

	// Create
	db.Create(&domain.Person{FirstName: "tom", LastName: "g", Username: "tom"})


	return db
}
