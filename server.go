package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func helloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello puttasiri",
	})
}

type Todo struct {
	ID     int    `json:"id"` //ขึ้นต้นด้วยตัวพิมพ์ใหญ่เท่านั้น
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "Active"},
}

func getTodosHandler(c echo.Context) error {
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/hello", helloHandler)
	//e.Start(":1323")

	e.GET("/todos", helloHandler)

	port := os.Getenv("PORT")
	log.Println("Port:", port)
	e.Start(":" + port)
}
