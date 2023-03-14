package data

import (
	modelFeedback "StayApp-API/features/feedback/data"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	UserID      int
	Name        string
	Price       int
	Description string
	Location    string
	UserName    string
	Pictures    []Pictures
	Feedback    []modelFeedback.Feedback
}

type Pictures struct {
	gorm.Model
	URL    string
	RoomID uint
}
