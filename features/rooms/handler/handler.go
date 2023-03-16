package handler

import (
	"StayApp-API/features/rooms"
	"StayApp-API/middlewares"
	"StayApp-API/utils/helper"
	"net/http"
	"strconv"

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
	addInput.UserID = uint(middlewares.ExtractToken(c))
	if err := c.Bind(&addInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	file, _ := c.FormFile("pictures")

	newRoom := rooms.Core{}
	copier.Copy(&newRoom, &addInput)

	err := rm.srv.Add(newRoom, file)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "add room successfully"))
}

func (rm *RoomHandler) GetAll(c echo.Context) error {
	var pageNumber int = 1
	pageParam := c.QueryParam("page")
	if pageParam != "" {
		pageConv, errConv := strconv.Atoi(pageParam)
		if errConv != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response("Failed, page must number"))
		} else {
			pageNumber = pageConv
		}
	}

	nameParam := c.QueryParam("name")
	data, err := rm.srv.GetAll(pageNumber, nameParam)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	}
	dataResponse := CoreToGetAllRoomResp(data)
	return c.JSON(http.StatusOK, helper.ResponseWithData("Success", dataResponse))
}

func (rm *RoomHandler) GetOne(c echo.Context) error {
	roomID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		return errCnv
	}
	data, err := rm.srv.GetOne(roomID)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := RoomResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "detail room successfully displayed", res))
}

func (rm *RoomHandler) Update(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	roomID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		return errCnv
	}
	// roomID := int(middlewares.ExtractToken(c))
	updateInput := RoomRequest{}
	if err := c.Bind(&updateInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	file, _ := c.FormFile("pictures")

	updateRoom := rooms.Core{}
	copier.Copy(&updateRoom, &updateInput)
	err := rm.srv.Update(userID, roomID, updateRoom, file)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusOK, "update room successfully"))
}

func (rm *RoomHandler) Delete(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	roomID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		return errCnv
	}
	err := rm.srv.Delete(userID, roomID)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusOK, "room successfully displayed"))
}

func (rm *RoomHandler) GetAllRoomUser(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	var pageNumber int = 1
	pageParam := c.QueryParam("page")
	if pageParam != "" {
		pageConv, errConv := strconv.Atoi(pageParam)
		if errConv != nil {
			return c.JSON(http.StatusInternalServerError, helper.Response("Failed, page must number"))
		} else {
			pageNumber = pageConv
		}
	}

	// nameParam := c.QueryParam("name")
	data, err := rm.srv.GetAllRoomUser(userID, pageNumber)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	}
	dataResponse := CoreToGetAllRoomByIdResp(data)
	return c.JSON(http.StatusOK, helper.ResponseWithData("Success", dataResponse))
}
