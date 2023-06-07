package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:255" json:"name"`
	Email       string `gorm:"size:50" json:"email"`
	Password    string `gorm:"size:255" json:"password"`
	PhoneNumber string `gorm:"size:50" json:"phoneNumber"`
	RoleId      string `gorm:"size:50" json:"role"`

	Role EnumType `gorm:"foreignKey:RoleId"`
}
