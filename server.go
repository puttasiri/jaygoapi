package main

import (
	"log"
	"net/http"
	"os"

	"github.com/puttasiri/jaygoapi/todo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("start middleware : check authentication")
		token := c.Request().Header.Get("Authorization")
		if token != "ABC" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}
		return next(c)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(authMiddleware)

	e.GET("/hello", todo.HelloHandler)
	//e.Start(":1323")

	e.GET("/todos", todo.GetTodosHandler)

	//insert
	e.POST("/todos", todo.CreateTodosHandler)

	//select by id
	e.GET("/todos/:id", todo.GetTodoByIdHandler)

	//update PUT
	e.PUT("/todos/:id", todo.UpdateTodosHandler)
	//delete DELETE
	e.DELETE("/todos/:id", todo.DeleteTodosHandler)

	port := os.Getenv("PORT")
	log.Println("Port:", port)
	e.Start(":" + port)
}
