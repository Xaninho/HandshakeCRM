package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	NIF                  int
	Name                 string
	CountryId            int
	StateId              int
	CurrencyId           int
	HasElectronicBilling bool
	PhotoURL             string
	PostalCode           string
	Address              string
	Notes                string

	Currency Currency `gorm:"foreignKey:CurrencyId"`
	Country  Country  `gorm:"foreignKey:CountryId"`
}
