package handler

import (
	"StayApp-API/features/feedback"
	"StayApp-API/middlewares"
	"StayApp-API/utils/helper"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type FeedbackHandler struct {
	srv feedback.FeedbackService
}

func New(service feedback.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{
		srv: service,
	}
}

func (fh *FeedbackHandler) Add(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	roomID, _ := strconv.Atoi(c.Param("id"))
	newFeedbackInput := NewFeedbackRequest{}
	if err := c.Bind(&newFeedbackInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	newFeedbackInput.UserID = uint(userID)
	newFeedbackInput.RoomID = uint(roomID)
	newFeedback := feedback.Core{}
	copier.Copy(&newFeedback, &newFeedbackInput)
	err := fh.srv.Add(newFeedback)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "added new feedback successfully"))
}
