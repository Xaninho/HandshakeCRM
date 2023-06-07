package models

import "gorm.io/gorm"

type EnumType struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey" json:"id"`
	Type        string `gorm:"size:255" json:"type"`
	Description string `gorm:"size:255" json:"description"`
	Active      bool   `gorm:"size:255" json:"active"`
}
