package router

import (
	_userData "StayApp-API/features/users/data"
	_userHandler "StayApp-API/features/users/handler"
	_userService "StayApp-API/features/users/services"
	_feedbackData "StayApp-API/features/feedback/data"
	_feedbackHandler "StayApp-API/features/feedback/handler"
	_feedbackService "StayApp-API/features/feedback/services"
	"StayApp-API/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userSrv := _userService.New(userData)
	userHdl := _userHandler.New(userSrv)
	e.POST("/login", userHdl.Login)
	e.POST("/register", userHdl.Register)
	e.GET("/users", userHdl.MyProfile, middlewares.JWTMiddleware())
	e.PUT("/users", userHdl.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", userHdl.Delete, middlewares.JWTMiddleware())
	e.PUT("/password", userHdl.ChangePassword, middlewares.JWTMiddleware())
	e.GET("/users/:id", userHdl.UserByID)

	feedbackData := _feedbackData.New(db)
	feedbackSrv := _feedbackService.New(feedbackData)
	feedbackHdl := _feedbackHandler.New(feedbackSrv)
	e.POST("/rooms/:id", feedbackHdl.Add, middlewares.JWTMiddleware())
}
