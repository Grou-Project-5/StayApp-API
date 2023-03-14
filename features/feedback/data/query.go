package data

import (
	"StayApp-API/features/feedback"

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