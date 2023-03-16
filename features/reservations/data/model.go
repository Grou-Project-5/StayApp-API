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
	GrossAmount   uint
}

func CoreToReservation(data reservations.Core) Reservation {
	return Reservation{
		Model:         gorm.Model{ID: data.ID},
		UserID:        data.UserID,
		RoomID:        data.RoomID,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		PaymentMethod: data.PaymentMethod,
		GrossAmount:   data.GrossAmount,
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
		GrossAmount:   data.GrossAmount,
	}
}
