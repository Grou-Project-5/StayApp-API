package data

import (
	"StayApp-API/features/feedback"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	RoomID    uint
	UserID    uint
	UserName  string
	Rating    int
	Comment   string
	CreatedAt string
}

func CoreToFeedback(data feedback.Core) Feedback {
	return Feedback{
		Model:     gorm.Model{ID: data.ID},
		RoomID:    data.RoomID,
		UserID:    data.UserID,
		UserName:  data.UserName,
		Rating:    data.Rating,
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt,
	}
}

func FeedbackToCore(data Feedback) feedback.Core {
	return feedback.Core{
		ID:        data.ID,
		RoomID:    data.RoomID,
		UserID:    data.UserID,
		UserName:  data.UserName,
		Rating:    data.Rating,
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt,
	}
}

func ListFeedbackToCore(data []Feedback) []feedback.Core {
	var dataCore []feedback.Core
	for _, v := range data {
		dataCore = append(dataCore, FeedbackToCore(v))
	}
	return dataCore
}
