package handler

import (
	"StayApp-API/features/rooms"
	"StayApp-API/utils/helper"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	srv rooms.RoomService
}

func New(srv rooms.RoomService) *RoomHandler {
	return &RoomHandler{
		srv: srv,
	}
}

func (rm *RoomHandler) Add(c echo.Context) error {
	addInput := RoomRequest{}
	if err := c.Bind(&addInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	newRoom := rooms.Core{}
	copier.Copy(&newRoom, &addInput)
	err := rm.srv.Add(newRoom)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "add room successfully"))
}
