package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID 			string 	`json:"id"`
	Isbn 		string 	`json:"isbn"`
	Title 	string 	`json:"title"`
	Author 	*Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string 	`json:"firstname"`
	Lastname string		`json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One",
		Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "825343", Title: "Book Two",
		Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "111111", Title: "Book Three",
		Author: &Author{Firstname: "Mike", Lastname: "Green"}})
	
	r.HandleFunc("/api/books", getBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
