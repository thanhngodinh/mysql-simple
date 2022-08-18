package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/thanhngodinh/mysql-simple/internal/route"
)

func main() {
	r := mux.NewRouter()
	route.RegisterBooksStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}