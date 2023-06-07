package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model

	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"size:255" json:"name"`
	CurrencyId int    `gorm:"size:50" json:"currencyId"`
}
