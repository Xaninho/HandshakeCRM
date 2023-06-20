package main

import (
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Company{},
		&models.Contact{},
		&models.Country{},
		&models.Currency{},
		&models.User{},
		&models.EnumType{},
		&models.Activity{},
	)
}
