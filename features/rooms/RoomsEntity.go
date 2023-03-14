package rooms

type Core struct {
	ID          int
	Name        string `validate:"required"`
	Price       int    `validate:"required"`
	Description string `validate:"required"`
	Location    string `validate:"required"`
	UserID      int
	UserName    string
	Pictures    []PictCore
}

type PictCore struct {
	ID     uint
	URL    string
	RoomID uint
}