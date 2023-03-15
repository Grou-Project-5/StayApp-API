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
func (rq *reservationQuery) Check(roomID int, startDate string, endDate string) (bool, error) {
	var tmp int64
	tx := rq.db.Model(&Reservation{}).Where("room_id = ?", roomID).Where("(start_date <= ? AND end_date >= ?) OR (start_date >= ? AND end_date <= ?) OR (start_date <= ? AND end_date >= ?)", endDate, startDate, startDate, endDate, startDate, endDate).Count(&tmp)
	if tx.Error != nil {
		return false, tx.Error
	}
	return tmp == 0, nil
}
