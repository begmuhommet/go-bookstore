package routes

import (
	"github.com/begmuhommet/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
