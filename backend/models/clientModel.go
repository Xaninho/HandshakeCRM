package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	Name            string
	Email           string
	PhotoURL        string
	PhoneNumber     int
	Aniversary      string
	CompanyId       int
	FunctionId      int
	TypeId          int
	AssignedAgentId int

	AssignedAgent Agent    `gorm:"foreignKey:AssignedAgentId"`
	Company       Company  `gorm:"foreignKey:CompanyId"`
	Type          EnumType `gorm:"foreignKey:TypeId"`
}
