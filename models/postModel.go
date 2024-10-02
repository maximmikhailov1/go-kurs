package models

import "gorm.io/gorm"

type Cars struct {
	gorm.Model
	Firm_name        string
	Model_name       string
	Reg_plate_number string
	Rent             int
	Is_detailed      bool
	Is_being_used    bool
}
