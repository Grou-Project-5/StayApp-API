package data

import (
	"StayApp-API/features/reservations"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID        uint
	RoomID        uint
	RoomName      string
	StartDate     string
	EndDate       string
	PaymentMethod string
	RoomPrice     int
	Days          int64
	GrossAmount   int64
	OrderID       string
	StatusPayment string `gorm:"default:Pending"`
	RedirectUrl   string
}

func CoreToReservation(data reservations.Core) Reservation {
	return Reservation{
		Model:         gorm.Model{ID: data.ID},
		UserID:        data.UserID,
		RoomID:        data.RoomID,
		RoomName:      data.RoomName,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		PaymentMethod: data.PaymentMethod,
		RoomPrice:     data.RoomPrice,
		Days:          data.Days,
		GrossAmount:   data.GrossAmount,
		OrderID:       data.OrderID,
		StatusPayment: data.StatusPayment,
		RedirectUrl:   data.RedirectUrl,
	}
}

func ReservationToCore(data Reservation) reservations.Core {
	return reservations.Core{
		ID:            data.ID,
		UserID:        data.UserID,
		RoomID:        data.RoomID,
		RoomName:      data.RoomName,
		StartDate:     data.StartDate,
		EndDate:       data.EndDate,
		PaymentMethod: data.PaymentMethod,
		RoomPrice:     data.RoomPrice,
		Days:          data.Days,
		GrossAmount:   data.GrossAmount,
		OrderID:       data.OrderID,
		StatusPayment: data.StatusPayment,
		RedirectUrl:   data.RedirectUrl,
	}
}

func ListReservationToCore(data []Reservation) []reservations.Core {
	var dataCore []reservations.Core
	for _, v := range data {
		dataCore = append(dataCore, ReservationToCore(v))
	}
	return dataCore
}