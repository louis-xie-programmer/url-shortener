package main

import (
	"url-shortener/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())

	// Routes
	e.GET("/", handler.IndexHandler)
	e.POST("/submit", handler.SubmitHandler)
	e.GET("/:id", handler.RedirectHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
