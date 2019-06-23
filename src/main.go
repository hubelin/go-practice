package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Got all books")
}
