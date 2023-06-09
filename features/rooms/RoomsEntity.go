package rooms

import "mime/multipart"

type Core struct {
	ID                uint
	Name              string `validate:"required"`
	Price             int    `validate:"required"`
	Description       string `validate:"required"`
	Location          string `validate:"required"`
	Availability      string
	UserID            uint
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
	UserPhone         string
	Pictures          string
	RatingRoom        float64
}

type RoomService interface {
	Add(newRoom Core, file *multipart.FileHeader) error
	GetAll(page int, name string) ([]Core, error)
	GetOne(id int) (Core, error)
	Update(userid int, id int, updateRoom Core, file *multipart.FileHeader) error
	Delete(userid int, id int) error
	GetAllRoomUser(userid int, page int) ([]Core, error)
}

type RoomData interface {
	Insert(input Core) error
	SelectAll(limit, offset int, name string) ([]Core, error)
	SelectOne(id uint) (Core, error)
	Update(userid uint, input Core, id uint) error
	Delete(userid int, id int) error
	SelectAllRoomUser(userid int, limit, offset int) ([]Core, error)
}
