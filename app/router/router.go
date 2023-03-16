package router

import (
	_userData "StayApp-API/features/users/data"
	_userHandler "StayApp-API/features/users/handler"
	_userService "StayApp-API/features/users/services"

	_roomData "StayApp-API/features/rooms/data"
	_roomHandler "StayApp-API/features/rooms/handler"
	_roomService "StayApp-API/features/rooms/services"

	_feedbackData "StayApp-API/features/feedback/data"
	_feedbackHandler "StayApp-API/features/feedback/handler"
	_feedbackService "StayApp-API/features/feedback/services"

	_reservationsData "StayApp-API/features/reservations/data"
	_reservationsHandler "StayApp-API/features/reservations/handler"
	_reservationsService "StayApp-API/features/reservations/services"
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

	roomData := _roomData.New(db)
	roomSrv := _roomService.New(roomData)
	roomHdl := _roomHandler.New(roomSrv)
	e.POST("/rooms", roomHdl.Add, middlewares.JWTMiddleware())
	e.GET("/rooms", roomHdl.GetAll)
	e.GET("/rooms/:id", roomHdl.GetOne)
	e.PUT("/rooms/:id", roomHdl.Update, middlewares.JWTMiddleware())
	e.DELETE("/rooms/:id", roomHdl.Delete, middlewares.JWTMiddleware())
	e.GET("/room", roomHdl.GetAllRoomUser, middlewares.JWTMiddleware())

	feedbackData := _feedbackData.New(db)
	feedbackSrv := _feedbackService.New(feedbackData)
	feedbackHdl := _feedbackHandler.New(feedbackSrv)
	e.POST("/feedbacks", feedbackHdl.Add, middlewares.JWTMiddleware())
	e.GET("rooms/:room_id/feedbacks", feedbackHdl.List)

	reservationsData := _reservationsData.New(db)
	reservationsSrv := _reservationsService.New(reservationsData)
	reservationsHdl := _reservationsHandler.New(reservationsSrv)
	e.POST("/check", reservationsHdl.Check)
	e.POST("/reservation", reservationsHdl.Add, middlewares.JWTMiddleware())
	e.POST("/paystatus", reservationsHdl.PayStatus)
}
