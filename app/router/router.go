package router

import (
	_userData "StayApp-API/features/users/data"
	_userHandler "StayApp-API/features/users/handler"
	_userService "StayApp-API/features/users/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userSrv := _userService.New(userData)
	userHdl := _userHandler.New(userSrv)
	e.POST("/login", userHdl.Login)
	e.POST("/register", userHdl.Register)
}
