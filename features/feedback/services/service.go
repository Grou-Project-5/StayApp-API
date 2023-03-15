package services

import (
	"StayApp-API/features/feedback"

	"github.com/go-playground/validator/v10"
)

type feedbackService struct {
	data feedback.FeedbackData
	vld  *validator.Validate
}

func New(repo feedback.FeedbackData) feedback.FeedbackService {
	return &feedbackService{
		data: repo,
		vld:  validator.New(),
	}
}

// Add implements feedback.FeedbackService
func (fs *feedbackService) Add(newFeedback feedback.Core) error {
	// Check input validation
	errVld := fs.vld.Struct(newFeedback)
	if errVld != nil {
		return errVld
	}
	err := fs.data.Add(newFeedback)
	if err != nil {
		return err
	}
	return nil
}

// List implements feedback.FeedbackService
func (fs *feedbackService) List(roomiD int) ([]feedback.Core, error) {
	tmp, err := fs.data.List(roomiD)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}