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
}
