package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	NIF                  int
	Name                 string
	CountryId            int
	CurrencyId           int
	HasElectronicBilling bool
	PostalCode           string
	Address              string
	Notes                string

	Currency Currency
	Country  Country
}
