package handler

type NewFeedbackRequest struct {
	RoomID  uint
	UserID  uint
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
