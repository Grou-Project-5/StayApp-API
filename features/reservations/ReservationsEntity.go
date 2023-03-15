package reservations

type Core struct {
	ID            uint
	UserID        uint
	RoomID        uint
	StartDate      string
	EndDate       string
	PaymentMethod string
}

type ReservationService interface {
	Check(roomID int, startDate, endDate string) (bool, error)
}

type ReservationData interface {
	Check(roomID int, startDate, endDate string) (bool, error)
}