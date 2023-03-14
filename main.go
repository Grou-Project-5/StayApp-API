package main

import (
	"StayApp-API/app/config"
	"StayApp-API/app/database"
	"StayApp-API/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)

	e := echo.New()
	e.Use(middleware.CORS())
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":80"))
}
