package routes

import (
	"github.com/RedrikShuhartRed/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBooksStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/book{Id}", controllers.DeleteBook).Methods("DELETE")

}
