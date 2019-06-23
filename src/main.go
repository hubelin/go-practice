package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Go1", Author: "Mr.Golang", Year: "2010"},
		Book{ID: 2, Title: "Go2", Author: "Mr.Golang", Year: "2020"},
		Book{ID: 3, Title: "Go3", Author: "Mr.Golang", Year: "2030"},
		Book{ID: 4, Title: "Go4", Author: "Mr.Golang", Year: "2040"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":8001", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// func NewEncoder returns a pointer to a new enconder that writes to w
	// func (*Encoder) Encode writes the JSON encoding of books to the stream
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	// func NewDecoder returns a pointer to a Decoder that reads from r.Body
	decoder := json.NewDecoder(r.Body)
	// func (*Decoder) Decode reads JSON encoded values from its input and stores it in v
	var book Book
	err := decoder.Decode(&book)

	if err != nil {
		log.Fatal(err)
	}

	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}
