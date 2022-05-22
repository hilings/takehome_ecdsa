package main

import (
	"myapp/handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.SetLevel(log.DEBUG)

	// routing
	e.GET("/get_message", handlers.GetMessage)
	e.POST("/verify", handlers.Verify)

	// server start
	e.Logger.Fatal(e.Start(":1323"))
}
