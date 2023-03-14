package data

import (
	modelFeedback "StayApp-API/features/feedback/data"
	modelRoom "StayApp-API/features/rooms/data"
	"StayApp-API/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Address  string
	Phone    string
	Gender   string `gorm:"type:enum('Male', 'Female')"`
	Pictures string
	Rooms    []modelRoom.Room
	Feedback []modelFeedback.Feedback
}

func CoreToUser(data users.Core) User {
	return User{
		Model:    gorm.Model{ID: data.Id},
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
		Phone:    data.Phone,
		Gender:   data.Gender,
		Pictures: data.Pictures,
	}
}

func UserToCore(data User) users.Core {
	return users.Core{
		Id:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Address:  data.Address,
		Phone:    data.Phone,
		Gender:   data.Gender,
		Pictures: data.Pictures,
	}
}