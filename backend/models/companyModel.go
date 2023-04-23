package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	NIF                 	int
    Name                	string
    Currency          	    Currency
    HasElectronicBilling	bool
    PhotoURL            	string
    Country           	    Country
    PostalCode          	string
    Address             	string
    State             	    EnumType
    Notes               	string
}