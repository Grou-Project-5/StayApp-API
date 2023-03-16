package services

import (
	"StayApp-API/features/rooms"
	"StayApp-API/utils/helper"
	"mime/multipart"

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

// Delete implements rooms.RoomService
func (srv *roomService) Delete(userid int, id int) error {
	errDelete := srv.data.Delete(userid, id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

// Update implements rooms.RoomService
func (srv *roomService) Update(userid int, id int, updateRoom rooms.Core, fileHeader *multipart.FileHeader) error {
	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.UploadFile(file, "/rooms")
		if err != nil {
			return err
		}
		updateRoom.Pictures = uploadURL[0]
	}
	errUpdate := srv.data.Update(uint(userid), updateRoom, uint(id))
	if errUpdate != nil {
		return errUpdate
	}
	return nil
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
func (srv *roomService) Add(newRoom rooms.Core, fileHeader *multipart.FileHeader) error {
	errValidate := srv.validate.Struct(newRoom)
	if errValidate != nil {
		return errValidate
	}
	if fileHeader != nil {
		file, _ := fileHeader.Open()
		uploadURL, err := helper.UploadFile(file, "/rooms")
		if err != nil {
			return err
		}
		newRoom.Pictures = uploadURL[0]
	}
	errInsert := srv.data.Insert(newRoom)
	if errInsert != nil {
		return errInsert
	}
	return nil
}
