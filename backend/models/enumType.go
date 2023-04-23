package models

import "gorm.io/gorm"

type EnumType struct {
	gorm.Model

	Type        	string
    Description 	string
    Active      	bool
}