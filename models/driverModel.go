package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Patronymic string
	Username   string `gorm:"uniqueIndex"`
	Password   string
	Car        *Car     `gorm:"constraint:OnDelete:SET NULL;"`
	CarID      *uint    `gorm:"constraint:OnDelete:SET NULL;"`
	Orders     []*Order `gorm:"foreignKey:DriverID"`
}
