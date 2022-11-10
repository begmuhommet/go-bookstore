package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/begmuhommet/go-bookstore/models"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook() {

}

func GetBookById() {

}

func UpdateBook() {

}

func DeleteBook() {

}
