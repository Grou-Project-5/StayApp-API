package services

import (
	"StayApp-API/features/rooms"

	"github.com/go-playground/validator/v10"
)

type roomService struct {
	data     rooms.RoomData
	validate *validator.Validate
}

func New(repo rooms.RoomData) rooms.RoomService {
	return &roomService{
		data:     repo,
		validate: validator.New(),
	}
}

// GetOne implements rooms.RoomService
func (srv *roomService) GetOne(id int) (rooms.Core, error) {
	data, err := srv.data.SelectOne(uint(id))
	return data, err
}

// GetAll implements rooms.RoomService
func (srv *roomService) GetAll(page int, name string) ([]rooms.Core, error) {
	limit := 10
	offset := (page - 1) * limit
	data, err := srv.data.SelectAll(limit, offset, name)
	return data, err
}

// Add implements rooms.RoomService
func (srv *roomService) Add(newRoom rooms.Core) error {
	errValidate := srv.validate.Struct(newRoom)
	if errValidate != nil {
		return errValidate
	}
	errInsert := srv.data.Insert(newRoom)
	if errInsert != nil {
		return errInsert
	}
	return nil
}
