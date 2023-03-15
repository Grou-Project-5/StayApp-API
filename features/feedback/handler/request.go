package handler

type NewFeedbackRequest struct {
	RoomID    uint
	UserID    uint
	CreatedAt string
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}
