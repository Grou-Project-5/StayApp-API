package feedback

type Core struct {
	ID        uint
	RoomID    uint
	UserID    uint
	UserName  string
	Rating    int    `validate:"required"`
	Comment   string `validate:"required"`
	CreatedAt string
}

type FeedbackService interface {
	Add(newFeedback Core) error
	List(roomiD int) ([]Core, error)
}

type FeedbackData interface {
	Add(newFeedback Core) error
	List(roomID int) ([]Core, error)
}
