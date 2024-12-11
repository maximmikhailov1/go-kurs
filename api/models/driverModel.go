package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	FirstName  string
	SecondName string
	Patronymic string
	Username   string `gorm:"uniqueIndex"`
	Password   string
	Car        *Car     `gorm:"constraint:OnDelete:SET NULL;"`
	CarID      *uint    `gorm:"constraint:OnDelete:SET NULL;"`
	Orders     []*Order `gorm:"foreignKey:DriverID"`
}

func (d *Driver) BeforeUpdate(tx *gorm.DB) (err error) {
	result := tx.Model(&Car{}).
		Where("id = ?", d.CarID).
		Update("is_being_used", false)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (d *Driver) AfterUpdate(tx *gorm.DB) (err error) {
	result := tx.Model(&Car{}).
		Where("id = ?", d.CarID).
		Update("is_being_used", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
