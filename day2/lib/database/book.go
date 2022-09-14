package database

import (
	"agmc/day2/models"
)

var books = []models.Book{
	{Id: 1, Judul: "buku 1", Tahun: 2019, Pengarang: "Bayu"},
	{Id: 2, Judul: "buku 2", Tahun: 2020, Pengarang: "Bayu"},
	{Id: 3, Judul: "buku 3", Tahun: 2021, Pengarang: "Bayu Test"},
}

func GetBooks() interface{} {
	return books
}

func GetOneBook(id int) interface{} {
	foundId := -1
	for i := 0; i < len(books); i++ {
		if books[i].Id == id {
			foundId = i
			break
		}
	}

	if foundId == -1 {
		return nil
	}

	return books[foundId]
}

func CreateBook(book models.Book) interface{} {
	books = append(books, book)
	return book
}

func UpdateBook(id int, book models.Book) interface{} {
	foundId := -1
	for i := 0; i < len(books); i++ {
		if books[i].Id == id {
			foundId = i
			break
		}
	}

	if foundId == -1 {
		return nil
	}
	books[foundId] = book
	return book
}

func DeleteBook(id int) interface{} {
	newBooks := []models.Book{}
	for i := 0; i < len(books); i++ {
		if books[i].Id != id {
			newBooks = append(newBooks, books[i])
		}
	}
	books = newBooks
	return nil
}
