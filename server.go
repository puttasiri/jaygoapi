package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello puttasiri",
	})
}
func main() {
	e := echo.New()
	e.GET("/hello", helloHandler)
	//e.Start(":1323")

	port := os.Getenv("PORT")
	log.Println("Port:", port)
	e.Start(":" + port)
}
