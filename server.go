package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello",
	})
}
func main() {
	e := echo.New()
	e.GET("/hello", helloHandler)
	e.Start(":1323")
}
