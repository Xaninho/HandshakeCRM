package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model

	ID            uint
	CompanyId     uint
	Name          string
	Email         string
	PhoneNumber   string
	Aniversary    string
	DispositionId int

	Company     Company
	Disposition EnumType
}
