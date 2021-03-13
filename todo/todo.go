package todo

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID     int    `json:"id"` //ขึ้นต้นด้วยตัวพิมพ์ใหญ่เท่านั้น
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay phone bills", Status: "Active"},
}

func HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello puttasiri",
	})
}

func GetTodosHandler(c echo.Context) error { //เปลี่ยน data เป็น ่json
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	return c.JSON(http.StatusOK, items)
}

func CreateTodosHandler(e echo.Context) error {
	t := Todo{}
	if err := e.Bind(&t); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t
	return e.JSON(http.StatusCreated, "created todo.")
}

// func GetTodoByIdHandler(c echo.Context) error {
// 	var id int
// 	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	t, ok := todos[id]
// 	if !ok {
// 		return c.JSON(http.StatusOK, map[int]string{})
// 	}
// 	return c.JSON(http.StatusOK, t)
// }
func UpdateTodosHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	t := todos[id]
	if err := c.Bind(t); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, t)
	// var id int
	// err := echo.PathParamsBinder(e).Int("id", &id).BindError()
	// if err != nil {
	// 	return e.JSON(http.StatusBadRequest, err)
	// }

	// t := Todo{}
	// if err := e.Bind(&t); err != nil {
	// 	return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	// }

	// id := len(todos)
	// id++
	// t.ID = id
	// todos[t.ID] = &t
	// return c.JSON(http.StatusOK, "updated todo.")
}

func DeleteTodosHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	delete(todos, id)
	return c.JSON(http.StatusOK, "deleted todo.")
}

func GetTodoByIdHandler(c echo.Context) error {
	var id int
	err := echo.PathParamsBinder(c).Int("id", &id).BindError()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos where id=$1") //select more row
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := id
	row := stmt.QueryRow(rowId)
	//var id int
	var title, status string

	err = row.Scan(&id, &title, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	//================================
	//
	// Todo struct {
	// 	ID     int    `json:"id"` //ขึ้นต้นด้วยตัวพิมพ์ใหญ่เท่านั้น
	// 	Title  string `json:"title"`
	// 	Status string `json:"status"`
	// }

	t := &Todo{
		ID:     id,
		Title:  title,
		Status: status,
	}
	// if !ok {
	// 	return c.JSON(http.StatusOK, map[int]string{})
	// }
	return c.JSON(http.StatusOK, t)

	//fmt.Println("one row", id, title, status)
}
