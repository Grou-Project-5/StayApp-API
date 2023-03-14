package feedback

type Core struct {
	Id       uint
	RoomID   uint
	UserID   uint
	UserName string
	Rating   int    `validate:"required"`
	Comment  string `validate:"required"`
}

type FeedbackService interface {
	Add(newFeedback Core) error
}

type FeedbackData interface {
	Add(newFeedback Core) error
}
