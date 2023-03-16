package services

import (
	"StayApp-API/features/reservations"

	"github.com/go-playground/validator/v10"
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
func (rs *reservationService) Check(roomID int, startDate string, endDate string) (bool, error) {
	tmp, err := rs.data.Check(roomID, startDate, endDate)
	if err != nil {
		return false, err
	}
	return tmp, nil
}


// Add implements reservations.ReservationService
func (rs *reservationService) Add(newReservation reservations.Core) error {
	// Check input validation
	errVld := rs.vld.Struct(newReservation)
	if errVld != nil {
		return errVld
	}
	err := rs.data.Add(newReservation)
	if err != nil {
		return err
	}
	return nil
}