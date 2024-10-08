package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name string
	Last_name  string
	Patronymic string
	Username   string
	Car        Car
}
