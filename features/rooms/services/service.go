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
