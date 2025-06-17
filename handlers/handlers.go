package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/001ajd/golang-books-api/models"
	"github.com/gorilla/mux"
)

func (h handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Books
	fmt.Println("Get all books API called ", r.RequestURI)
	h.DB.Find(&books)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	vars := mux.Vars(r)
	bookId := vars["id"]
	h.DB.Where("id = ?", bookId).Find(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
