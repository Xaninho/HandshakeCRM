package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model

	Code         string
	Name         string
	ExchangeRate float64
}
