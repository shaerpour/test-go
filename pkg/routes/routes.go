package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json: "id"`
	Name   string `json: "name"`
	Author string `json: "author"`
}

var BookList = []Book{
	{ID: 1, Name: "ahsp", Author: "ahsp"},
	{ID: 2, Name: "parmin", Author: "parmin"},
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(BookList)
}

func GetBooksByID(w http.ResponseWriter, r *http.Request) {
	book_id, _ := strconv.Atoi(mux.Vars(r)["id"])
	for index, book := range BookList {
		if book.ID == book_id {
			json.NewEncoder(w).Encode(BookList[index])
		}
	}
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	var NewBook Book

	json.NewDecoder(r.Body).Decode(&NewBook)
	for _, book := range BookList {
		if book.ID == NewBook.ID || book.Author == NewBook.Author && book.Name == NewBook.Name {
			json.NewEncoder(w).Encode(http.StatusConflict)
			return
		}
	}
	BookList = append(BookList, NewBook)
	json.NewEncoder(w).Encode(BookList)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	book_id, _ := strconv.Atoi(mux.Vars(r)["id"])

	for index, book := range BookList {
		if book.ID == book_id {
			BookList = append(BookList[0:index], BookList[index+1:]...)
			json.NewEncoder(w).Encode(BookList)
		}
	}
}
