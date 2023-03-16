package reservations

import "github.com/midtrans/midtrans-go/snap"

type Core struct {
	ID            uint
	UserID        uint
	RoomID        uint   `validate:"required"`
	StartDate     string `validate:"required"`
	EndDate       string `validate:"required"`
	PaymentMethod string
	RoomPrice     int
	Days          int64
	GrossAmount   int64
	OrderID       string
	StatusPaymen  string
}

type ReservationService interface {
	Check(checkAvailability Core) (bool, error)
	Add(newReservation Core) (string, *snap.Response, error)
}

type ReservationData interface {
	Check(checkAvailability Core) (bool, error)
	GrossAmt(roomID int, days int64) int64
	Add(newReservation Core) error
}
