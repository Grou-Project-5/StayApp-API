package handler

type ReservationRequest struct {
	RoomID    uint `json:"room_id"`
	UserID    uint
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	OrderID   string
	Days      int64
}

type PayStatusRequest struct {
	OrderID   string `json:"order_id"`
}
