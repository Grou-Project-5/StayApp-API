package handler

import (
	"StayApp-API/features/reservations"
	"StayApp-API/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	srv reservations.ReservationService
}

func New(srv reservations.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		srv: srv,
	}
}

func (rm *ReservationHandler) Check(c echo.Context) error {
	checkInput := CheckRequest{}
	if err := c.Bind(&checkInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	tmp, err := rm.srv.Check(int(checkInput.RoomID), checkInput.StartDate, checkInput.EndDate)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := Availability{}
	res.Availability = tmp
	return c.JSON(helper.SuccessResponse(http.StatusOK, "room available", res))
}
