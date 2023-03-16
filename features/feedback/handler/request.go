package handler

type NewFeedbackRequest struct {
	RoomID    uint	`json:"room_id"`
	UserID    uint
	CreatedAt string
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}
