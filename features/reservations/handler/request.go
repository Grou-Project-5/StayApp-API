package handler

type CheckRequest struct {
	RoomID   uint   `json:"name"`
	StartDate string `json:"start_date"`
	EndDate  string `json:"end_date"`
}
