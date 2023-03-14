package handler

import (
	"StayApp-API/features/users"
	"StayApp-API/middlewares"
	"strconv"

	// "StayApp-API/middlewares"
	"StayApp-API/utils/helper"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	srv users.UserService
}

func New(service users.UserService) *UserHandler {
	return &UserHandler{
		srv: service,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	registerInput := RegisterRequest{}
	if err := c.Bind(&registerInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	newUser := users.Core{}
	copier.Copy(&newUser, &registerInput)
	err := uh.srv.Register(newUser)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "register successfully"))
}

func (uh *UserHandler) Login(c echo.Context) error {
	loginInput := LoginRequest{}
	if err := c.Bind(&loginInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	token, data, err := uh.srv.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := MiniResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "login successfully", res, token))
}

func (uh *UserHandler) MyProfile(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	data, err := uh.srv.MyProfile(userID)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := UserResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, "profile successfully displayed", res))
}

func (uh *UserHandler) Update(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	updateInput := UpdateRequest{}
	if err := c.Bind(&updateInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	file, errFile := c.FormFile("pictures")
	if errFile != nil {
		return errFile
	}
	updateUser := users.Core{}
	copier.Copy(&updateUser, &updateInput)
	err := uh.srv.Update(userID, updateUser, file)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusOK, "update profile successfully"))
}

func (uh *UserHandler) ChangePassword(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c))
	passwordInput := ChangePasswordRequest{}
	if err := c.Bind(&passwordInput); err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	updatePassword := users.Core{}
	updatePassword.Password = passwordInput.NewPass
	err := uh.srv.ChangePassword(userID, passwordInput.OldPass, updatePassword)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusOK, "changed password successfully"))
}

func (uh *UserHandler) UserByID(c echo.Context) error {
	userID, errCnv := strconv.Atoi(c.Param("id"))
	if errCnv != nil {
		return errCnv
	}
	data, err := uh.srv.UserByID(userID)
	if err != nil {
		return c.JSON(helper.ErrorResponse(err))
	}
	res := UserResponse{}
	copier.Copy(&res, &data)
	return c.JSON(helper.SuccessResponse(http.StatusOK, " profile successfully displayed", res))
}
