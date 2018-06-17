package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)



func Bootstrap() * gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Person{})

	// Create
	db.Create(&Person{FirstName: "tom", LastName: "g", Username: "tom"})

	// Read
	var person Person
	db.First(&person, 1)                      // find product with id 1
	// Update - update product's price to 2000
	db.Model(&person).Update("LastName", "gg")

	// Delete - delete product
	db.Delete(&person)
	return db
}
