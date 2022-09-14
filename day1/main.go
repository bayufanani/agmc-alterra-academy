package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string
	Email string
}

func main() {
	e := echo.New()

	e.GET("/", HelloController)
	e.GET("/user", GetUser)

	e.Start(":8080")
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "hellow world")
}

func GetUser(c echo.Context) error {
	user := User{Name: "Bayu Fanani", Email: "fanani707@gmail.com"}
	return c.JSON(http.StatusOK, user)
}
