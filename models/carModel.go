package models

type Car struct {
	ID               uint `gorm:"primaryKey"`
	Firm_name        string
	Model_name       string
	Reg_plate_number string
	VIN_number       string
	Rent             int
	Is_detailed      bool
	Is_being_used    bool
	User             User `gorm:"foreignKey:CarID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
