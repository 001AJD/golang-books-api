package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Books struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

var books = []Books{
	{Name: "Atomis Habits", Author: "James Clear"},
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all books API called ", r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
