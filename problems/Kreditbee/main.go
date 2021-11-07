package main

import (
	"Kreditbee/db_create"
	"Kreditbee/handler"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loding env")
	}
	e := echo.New()
	handler.SetRoutes(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var photos handler.Photos
	var albums handler.Albums
	db_create.GetDb_instance().AutoMigrate(&photos)
	db_create.GetDb_instance().AutoMigrate(&albums)
	// Middleware
	e.Start(":1200")
}
