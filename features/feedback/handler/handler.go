package handler

import (
	"StayApp-API/features/feedback"
	"StayApp-API/middlewares"
	"StayApp-API/utils/helper"
	"log"
	"net/http"
	"strconv"
	"time"

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
	curtime := time.Now()
	newFeedbackInput.CreatedAt = curtime.Format("January 2006")
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

func (fh *FeedbackHandler) List(c echo.Context) error {
	roomID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		return errCnv
	}
	data, err := fh.srv.List(roomID)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := ListFeedbackResponse{}
	copier.Copy(&res, &data)
	log.Println(res)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "displayed feedback successfully", res))
}
