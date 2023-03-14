package feedback

type Core struct {
	Id      int
	RoomID  int
	UserID  int
	Rating  int    `validate:"required"`
	Comment string `validate:"required"`
}
