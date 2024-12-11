package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Username   string
	Password   string
	FirstName  string
	SecondName string
	Orders     []Order `gorm:"foreignKey:ClientID"`
}
