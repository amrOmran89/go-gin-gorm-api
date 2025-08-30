package initializers

import (
	models "go-gin-gorm/entities"
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

	DB.AutoMigrate(&models.PostEntity{})
	DB.AutoMigrate(&models.UserEntity{})
	log.Printf("Database connected successfully")
}
