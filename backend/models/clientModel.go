package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	Company  		Company
    Name       		string
    Email      		string
    PhotoURL   		string
    Function 		EnumType
    Type     		EnumType
    PhoneNumber 	int
    Aniversary 		string
}