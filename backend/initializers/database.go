package initializers

import (
	"log"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	/*var err error

	server := os.Getenv("AZURE_SERVER")
	user := os.Getenv("AZURE_USER")
	password := os.Getenv("AZURE_PASSWORD")
	port := os.Getenv("AZURE_PORT")
	database := os.Getenv("AZURE_DATABASE")

	dsn := "server=" + server + ";user id=" + user + ";password=" + password + ";port=" + port + ";database=" + database + ";"
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect")
		log.Fatal(DB)
	}*/
	var err error
	dsn := os.Getenv("DB_CONNECTION_STRING")
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect")
		log.Fatal(DB)
	}
}
