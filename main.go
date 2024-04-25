package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaerpour/test-go/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/books/all", routes.GetAllBooks).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", routes.GetBooksByID).Methods("GET")
	r.HandleFunc("/api/v1/books/add", routes.AddNewBook).Methods("PUT")
	r.HandleFunc("/api/v1/books/delete/{id}", routes.DeleteBook).Methods("DELETE")

	fmt.Println("Running Server on port 8080")
	http.ListenAndServe(":8080", r)
}
