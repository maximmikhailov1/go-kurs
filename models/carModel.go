package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Firm_name        string
	Model_name       string
	Reg_plate_number string
	VIN_number       string `gorm:"size:17"`
	Rent             int
	Is_detailed      bool
	Is_being_used    bool
}
