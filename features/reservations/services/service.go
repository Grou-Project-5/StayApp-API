package services

import (
	"StayApp-API/features/reservations"
	"StayApp-API/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/midtrans/midtrans-go/snap"
)

type reservationService struct {
	data reservations.ReservationData
	vld  *validator.Validate
}

func New(repo reservations.ReservationData) reservations.ReservationService {
	return &reservationService{
		data: repo,
		vld:  validator.New(),
	}
}

// Check implements reservations.ReservationService
func (rs *reservationService) Check(checkAvailability reservations.Core) (bool, error) {
	tmp, err := rs.data.Check(checkAvailability)
	if err != nil {
		return false, err
	}
	return tmp, nil
}

// Add implements reservations.ReservationService
func (rs *reservationService) Add(newReservation reservations.Core) (string, *snap.Response, error) {
	// Check input validation
	errVld := rs.vld.Struct(newReservation)
	if errVld != nil {
		return "", nil, errVld
	}
	GrossAmount := rs.data.GrossAmt(int(newReservation.RoomID), newReservation.Days)

	midtrans := helper.MidtransPay(newReservation.OrderID, GrossAmount)
	newReservation.GrossAmount = GrossAmount
	err := rs.data.Add(newReservation)
	if err != nil {
		return "", nil, err
	}
	return newReservation.OrderID, midtrans, nil
}
