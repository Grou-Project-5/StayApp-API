package data

import (
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID        uint
	RoomID        uint
	StartDate      string
	EndDate       string
	PaymentMethod string
}
