package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model

	Name        	string
    Email       	string
    Position  	    EnumType
    PhoneNumber 	int
    PhotoURL    	string
}