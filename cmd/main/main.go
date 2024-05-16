package main

import (
	"log"
	"net/http"

	"github.com/RedrikShuhartRed/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
