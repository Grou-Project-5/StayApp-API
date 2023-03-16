package handler

type NewFeedbackRequest struct {
	RoomID    uint `json:"room_id"`
	UserID    uint
	CreatedAt string
	Rating    float64 `json:"rating"`
	Comment   string  `json:"comment"`
}
