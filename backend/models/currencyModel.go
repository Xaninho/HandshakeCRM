package models

import "gorm.io/gorm"

type Currency struct {
	gorm.Model

	ID           uint    `gorm:"primaryKey" json:"id"`
	Code         string  `gorm:"size:255" json:"code"`
	Name         string  `gorm:"size:255" json:"name"`
	ExchangeRate float64 `gorm:"size:255" json:"exchangeRate"`
}
