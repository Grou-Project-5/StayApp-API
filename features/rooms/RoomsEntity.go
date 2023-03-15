package rooms

import "mime/multipart"

type Core struct {
	ID                uint
	Name              string `validate:"required"`
	Price             int    `validate:"required"`
	Description       string `validate:"required"`
	Location          string `validate:"required"`
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
	Pictures          []PictCore
}

type PictCore struct {
	ID     uint
	URL    string
	RoomID uint
}

type RoomService interface {
	Add(newRoom Core) error
	GetAll(page int, name string) ([]Core, error)
	GetOne(id int) (Core, error)
	Update(id int, updateRoom Core, file *multipart.FileHeader) error
	Delete(id int) error
}

type RoomData interface {
	Insert(input Core) error
	SelectAll(limit, offset int, name string) ([]Core, error)
	SelectOne(id uint) (Core, error)
	Update(input Core, id uint) error
	Delete(id int) error
}
