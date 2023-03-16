package data

import (
	"StayApp-API/features/reservations"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID        uint
	RoomID        uint
	StartDate     string
	EndDate       string
	PaymentMethod string
	RoomPrice     int
	Days          int64
	GrossAmount   int64
	OrderID       string
	StatusPaymen  string `gorm:"default:Pending"`
}

func CoreToReservation(data reservations.Core) Reservation {
	return Reservation{
		Model:         gorm.Model{ID: data.ID},
		UserID:        data.UserID,
		RoomID:        data.RoomID,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		PaymentMethod: data.PaymentMethod,
		RoomPrice:     data.RoomPrice,
		Days:          data.Days,
		GrossAmount:   data.GrossAmount,
		OrderID:       data.OrderID,
		StatusPaymen:  data.StatusPaymen,
	}
}

func ReservationToCore(data Reservation) reservations.Core {
	return reservations.Core{
		ID:            data.ID,
		UserID:        data.UserID,
		RoomID:        data.RoomID,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		PaymentMethod: data.PaymentMethod,
		RoomPrice:     data.RoomPrice,
		Days:          data.Days,
		GrossAmount:   data.GrossAmount,
		OrderID:       data.OrderID,
		StatusPaymen:  data.StatusPaymen,
	}
}
