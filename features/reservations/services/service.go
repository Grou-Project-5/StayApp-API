package services

import (
	"StayApp-API/features/reservations"
	"StayApp-API/utils/helper"
	"errors"

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
	// cek ketersediaan kamar
	tmp, _ := rs.data.Check(newReservation)
	if !tmp {
		return "", nil, errors.New("room not available")
	}
	// hitung total pembayaran
	grossAmount := rs.data.GrossAmt(int(newReservation.RoomID), newReservation.Days)
	// cek owner room
	owner := rs.data.CheckOwner(int(newReservation.RoomID), int(newReservation.UserID))
	if !owner {
		return "", nil, errors.New("tidak bisa memesan kamar milik sendiri")
	}

	midtrans := helper.MidtransPay(newReservation.OrderID, grossAmount)
	newReservation.GrossAmount = grossAmount
	newReservation.RedirectUrl = midtrans.RedirectURL
	err := rs.data.Add(newReservation)
	if err != nil {
		return "", nil, err
	}
	return newReservation.OrderID, midtrans, nil
}

// PayStatus implements reservations.ReservationService
func (rs *reservationService) PayStatus(updatePayStatus reservations.Core) (reservations.Core, error) {
	data, errStatus := helper.MidtransStatusPay(updatePayStatus.OrderID)
	if errStatus != nil {
		return reservations.Core{}, errStatus
	}
	if data.TransactionStatus == "settlement" {
		updatePayStatus.StatusPayment = "success"
	}
	err := rs.data.PayStatus(updatePayStatus)
	if err != nil {
		return reservations.Core{}, err
	}
	return updatePayStatus, nil
}
