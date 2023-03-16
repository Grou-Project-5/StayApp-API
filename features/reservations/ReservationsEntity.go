package reservations

type Core struct {
	ID            uint
	UserID        uint
	RoomID        uint   `validate:"required"`
	StartDate     string `validate:"required"`
	EndDate       string `validate:"required"`
	PaymentMethod string
	GrossAmount   uint
}

type ReservationService interface {
	Check(roomID int, startDate, endDate string) (bool, error)
	Add(newReservation Core) error
}

type ReservationData interface {
	Check(roomID int, startDate, endDate string) (bool, error)
	Add(newReservation Core) error
}
