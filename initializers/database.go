package initializers

import (
	"go-gin-gorm/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dbName := GetDatabaseName()
	log.Printf("Connecting to database: %s", dbName)

	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})
	log.Printf("Database connected successfully")
}
