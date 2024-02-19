package routes

import (
	"golang-book-management-system/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var ReigisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books", controllers.ListBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods(http.MethodDelete)
}
