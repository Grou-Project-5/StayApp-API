package data

import (
	modelRoom "StayApp-API/features/rooms/data"
	modelFeedback "StayApp-API/features/feedback/data"

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
