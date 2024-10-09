package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	First_name string
	Last_name  string
	Patronymic string
	Username   string `gorm:"uniqueIndex"`
	CarID      uint   `gorm:"default:7"`
}
