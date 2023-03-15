package data

import (
	"StayApp-API/features/feedback"
	"errors"

	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.FeedbackData {
	return &feedbackQuery{
		db: db,
	}
}

// Add implements feedback.FeedbackData
func (fq *feedbackQuery) Add(newFeedback feedback.Core) error {
	data := CoreToFeedback(newFeedback)
	tx := fq.db.Create(&data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// List implements feedback.FeedbackData
func (fq *feedbackQuery) List(roomID int) ([]feedback.Core, error) {
	tmp := []Feedback{}
	tx := fq.db.Where("room_id = ?", roomID).Select("feedbacks.id, feedbacks.rating, feedbacks.comment, feedbacks.created_at, users.name AS user_name").Joins("JOIN users ON feedbacks.user_id = users.id").Find(&tmp)
	if tx.RowsAffected < 1 {
		return nil, errors.New("no comments yet")
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	listFeedback := ListFeedbackToCore(tmp)
	return listFeedback, nil
}
