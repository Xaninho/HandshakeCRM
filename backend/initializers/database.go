package initializers

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB * gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB_CONNECTION_STRING")
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect")
		log.Fatal(DB)
	}
}