package routes

import (
	"agmc/day2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/books", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetOneBookController)
	e.POST("/books", controllers.CreateBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)

	e.GET("/users", controllers.GetUserControllers)
	e.GET("/users/:id", controllers.GetOneUserController)
	e.POST("/users", controllers.CreateUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	return e
}
