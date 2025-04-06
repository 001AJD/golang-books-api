package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/001ajd/golang-books-api/models"
)

func (h handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Books
	fmt.Println("Get all books API called ", r.RequestURI)
	h.DB.Find(&books)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
