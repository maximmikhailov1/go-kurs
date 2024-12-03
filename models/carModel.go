package models

type Car struct {
	ID             uint `gorm:"primaryKey"`
	FirmName       string
	ModelName      string
	RegPlateNumber string
	VinNumber      string
	Rent           int
	Color          string
	IsDetailed     bool
	IsBeingUsed    bool
}
