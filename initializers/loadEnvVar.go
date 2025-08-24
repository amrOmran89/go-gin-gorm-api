package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing without environment variables")
	}
}

func GetEnvVar(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetDatabaseName() string {
	env := GetEnvVar("ENV", "development")

	switch env {
	case "development":
		return GetEnvVar("TEST_DB_NAME", "test.db")
	case "production":
		return GetEnvVar("DB_NAME", "prod.db")
	default:
		return GetEnvVar("TEST_DB_NAME", "test.db")
	}
}
