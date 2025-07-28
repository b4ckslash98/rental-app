package entity

type Order struct {
	ID              int `gorm:"primaryKey"`
	CarID           int
	UserID          int
	OrderDate       string
	PickupDate      string
	DropoffDate     string
	PickupLocation  string
	DropoffLocation string
}
