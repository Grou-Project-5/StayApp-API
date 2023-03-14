package rooms

type Core struct {
	ID                int
	Name              string `validate:"required"`
	Price             int    `validate:"required"`
	Description       string `validate:"required"`
	Location          string `validate:"required"`
	UserID            int
	UserName          string
	MaxVisitors       int
	SpecialAccess     string
	Bedroom           int
	RoomTotal         int
	Kitchen           string
	Wifi              string
	Garage            string
	ExcellentFeatures string
	CleaningFee       int
	ServiceFee        int
	BringAnimal       string
	CheckIn           string
	CheckOut          string
	TakePhoto         string
	Other             string
	Pictures          []PictCore
}

type PictCore struct {
	ID     uint
	URL    string
	RoomID uint
}

type RoomService interface {
	Add(newRoom Core) error
}

type RoomData interface {
	Insert(input Core) error
}
