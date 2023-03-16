package handler

import (
	"StayApp-API/features/reservations"
	"StayApp-API/middlewares"
	"StayApp-API/utils/helper"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
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
	checkInput := ReservationRequest{}
	if err := c.Bind(&checkInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}

	checkAvailability := reservations.Core{}
	copier.Copy(&checkAvailability, &checkInput)

	tmp, err := rm.srv.Check(checkAvailability)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := Availability{}
	var message string
	res.Availability = tmp
	if !tmp {
		message = "room not available"
	} else {
		message = "room available"
	}
	return c.JSON(helper.SuccessResponse(http.StatusOK, message, res))
}

func (rm *ReservationHandler) Add(c echo.Context) error {
	newReservationInput := ReservationRequest{}
	if err := c.Bind(&newReservationInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}

	starDate, _ := time.Parse("2006-01-02", newReservationInput.StartDate)
	endDate, _ := time.Parse("2006-01-02", newReservationInput.EndDate)
	newReservationInput.Days = int64((endDate.Sub(starDate).Hours() / 24) + 1)
	
	now := time.Now()
	unixTimeNano := now.UnixNano()
	newReservationInput.OrderID = "HAPP-" + strconv.FormatInt(unixTimeNano, 10)
	newReservationInput.UserID = uint(middlewares.ExtractToken(c))

	newReservation := reservations.Core{}
	copier.Copy(&newReservation, &newReservationInput)
	orderID, data, err := rm.srv.Add(newReservation)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := PaymentResponse{}
	res.OrderID = orderID
	res.RedirectURL = data.RedirectURL
	return c.JSON(helper.SuccessResponse(http.StatusOK, "reservation success, waiting for payment", res))
}
