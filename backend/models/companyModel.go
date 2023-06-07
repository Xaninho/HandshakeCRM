package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	ID                   uint   `gorm:"primaryKey" json:"id"`
	NIF                  int    `gorm:"size:255" json:"nif"`
	Name                 string `gorm:"size:255" json:"name"`
	CountryId            int    `gorm:"size:50" json:"countryId"`
	CurrencyId           int    `gorm:"size:50" json:"currencyId"`
	HasElectronicBilling bool   `gorm:"size:255" json:"hasElectronicBilling"`
	PostalCode           string `gorm:"size:255" json:"postalCode"`
	Address              string `gorm:"size:255" json:"address"`
	Notes                string `gorm:"size:255" json:"notes"`

	Currency Currency `gorm:"foreignKey:CurrencyId"`
	Country  Country  `gorm:"foreignKey:CountryId"`
}
