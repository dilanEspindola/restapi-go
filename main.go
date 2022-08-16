package main

import (
	"github.com/gorilla/mux"

	"encoding/json"

	"math/rand"
	"net/http"
	"strconv"
)

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// init books
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

// get book
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // get param
	for _, book := range books {
		if book.Id == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Book{})
}

// create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	// fake data
	books = append(books, Book{
		Id:    "1",
		Isbn:  "44567",
		Title: "book one",
		Author: &Author{
			FirstName: "Andrew",
			LastName:  "Doe",
		},
	})
	books = append(books, Book{
		Id:    "2",
		Isbn:  "44569",
		Title: "book two",
		Author: &Author{
			FirstName: "Andres",
			LastName:  "Perez",
		},
	})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":4000", r)
}
