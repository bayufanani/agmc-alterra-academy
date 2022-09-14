package controllers

import (
	"agmc/day2/lib/database"
	"agmc/day2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBookController(c echo.Context) error {
	books := database.GetBooks()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"buku":   books,
	})
}

func GetOneBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := database.GetOneBook(id)
	if book == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"buku":   book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	b := database.CreateBook(book)
	if b == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "tidak bisa menambahkan buku")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"buku":   b,
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	c.Bind(&book)
	b := database.UpdateBook(id, book)
	if b == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "tidak ditemukan")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"buku":   b,
	})
}

func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DeleteBook(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
