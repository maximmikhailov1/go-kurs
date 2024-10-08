package models

type Car struct {
	ID               int
	Firm_name        string
	Model_name       string
	Reg_plate_number string
	VIN_number       string
	Rent             int
	Is_detailed      bool
	Is_being_used    bool
}
