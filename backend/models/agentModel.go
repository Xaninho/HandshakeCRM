package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model

	Name        string
	Email       string `gorm:"unique"`
	Password    string
	PositionId  int
	PhoneNumber int
	PhotoURL    string

	Position EnumType `gorm:"foreignKey:PositionId"`
}
