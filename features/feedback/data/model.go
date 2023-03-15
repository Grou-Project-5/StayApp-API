package data

import (
	"StayApp-API/features/feedback"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	RoomID   uint
	UserID   uint
	UserName string
	Rating   int
	Comment  string
}

func CoreToFeedback(data feedback.Core) Feedback {
	return Feedback{
		Model:    gorm.Model{ID: data.Id},
		RoomID:   data.RoomID,
		UserID:   data.UserID,
		UserName: data.UserName,
		Rating:   data.Rating,
		Comment:  data.Comment,
	}
}
