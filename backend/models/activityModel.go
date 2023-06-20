package models

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model

	ContactID uint
	UserID    uint

	Date      string
	Title     string
	Location  string
	Duration  string
	Notes     string
	OutcomeId int

	Contact Contact
	User    User
	Outcome EnumType
}
