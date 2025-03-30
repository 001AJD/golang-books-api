package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/001ajd/golang-books-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("books-api running...")
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/books", handlers.GetBooks).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))

}
