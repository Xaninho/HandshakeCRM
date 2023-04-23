package main

import (
	"github.com/handshakeCRM/initializers"
	"github.com/handshakeCRM/models"
)


func init() {
	initializers.LoadEnvVariables();
	initializers.ConnectToDb();
}

func main() {
	initializers.DB.AutoMigrate(
        &models.Company{},
        &models.Client{},
        &models.Country{},
        &models.Currency{},
        &models.Agent{},
        &models.EnumType{},
    )
}