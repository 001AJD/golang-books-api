package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/001ajd/golang-books-api/db"
	"github.com/001ajd/golang-books-api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load the .env file")
	}
	fmt.Println("books-api running...")
	DB := db.ConnectDb()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/books", h.GetBooks).Methods("GET")
	router.HandleFunc("/api/v1/books/{id}", h.GetBookById).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))

}
