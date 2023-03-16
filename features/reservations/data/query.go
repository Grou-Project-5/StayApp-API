package data

import (
	"StayApp-API/features/reservations"

	"gorm.io/gorm"
)

type reservationQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservations.ReservationData {
	return &reservationQuery{
		db: db,
	}
}

// Check implements reservations.ReservationData
func (rq *reservationQuery) Check(checkAvailability reservations.Core) (bool, error) {
	data := CoreToReservation(checkAvailability)
	var tmp int64
	tx := rq.db.Model(&Reservation{}).Where("room_id = ?", data.RoomID).Where("(start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?) OR (start_date <= ? AND end_date >= ?)", data.EndDate, data.StartDate, data.StartDate, data.EndDate, data.StartDate, data.EndDate).Count(&tmp)
	if tx.Error != nil {
		return false, tx.Error
	}
	return tmp == 0, nil
}

// GrossAmt implements reservations.ReservationData
func (rq *reservationQuery) GrossAmt(roomID int, days int64) int64 {
	var totalPrice int64
	rq.db.Raw("SELECT price FROM rooms WHERE id = ?", roomID).First(&totalPrice)
	totalPrice = int64(days) * totalPrice
	return totalPrice
}

// CheckOwner implements reservations.ReservationData
func (rq *reservationQuery) CheckOwner(roomID int, userID int) bool {
	var tmp int64
	rq.db.Raw("SELECT COUNT(*) FROM rooms WHERE id = ? AND user_id = ?", roomID, userID).Scan(&tmp)
	return tmp == 0
}

// Add implements reservations.ReservationData
func (rq *reservationQuery) Add(newReservation reservations.Core) error {
	data := CoreToReservation(newReservation)
	tx := rq.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// PayStatus implements reservations.ReservationData
func (rq *reservationQuery) PayStatus(updatePayStatus reservations.Core) error {
	data := CoreToReservation(updatePayStatus)
	tx := rq.db.Model(&Reservation{}).Where("order_id = ?", updatePayStatus.OrderID).Updates(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
