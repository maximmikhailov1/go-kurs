package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	DriverID     uint
	ClientID     uint
	AddressFrom  string
	AddressWhere string
}
