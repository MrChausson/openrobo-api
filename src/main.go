package main

import (
	"net/http"
	"openrobo-api/src/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
	e.GET("/introduce", handler.IntroduceHandler)
	e.GET("/ask", handler.AskHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
