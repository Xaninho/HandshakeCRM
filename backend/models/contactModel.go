package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model

	ID            uint   `gorm:"primaryKey" json:"id"`
	CompanyId     uint   `json:"companyId"`
	Name          string `gorm:"size:255" json:"name"`
	Email         string `gorm:"size:255" json:"email"`
	PhoneNumber   string `gorm:"size:20" json:"phoneNumber"`
	Aniversary    string `gorm:"size:20" json:"aniversary"`
	DispositionId int    `json:"dispositionId"`

	Company     Company  `gorm:"foreignKey:CompanyId"`
	Disposition EnumType `gorm:"foreignKey:DispositionId"`
}
