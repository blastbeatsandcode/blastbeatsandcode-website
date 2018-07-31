package utils

import (
	"fmt"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/* GetDB returns a connection to the database */
func GetDB() *gorm.DB {
	// Connect to the database, it follows this patter
	// gorm.Open("mysql", username:password@(localhost:port)/databasename)
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/blastbeats")

	if err != nil {
		fmt.Println("Error getting database connection")
		fmt.Println(err)
	}

	dbSetup(db)

	return db
}

/* Automigrate database based on the structs in /models */
func dbSetup(db *gorm.DB) gorm.DB {
	db.AutoMigrate(&models.BlogPost{})
	db.AutoMigrate(&models.User{})

	return *db
}
