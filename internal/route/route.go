package route

import (
	"github.com/gorilla/mux"
	"github.com/thanhngodinh/mysql-simple/internal/controller"
)

var Router = func(router *mux.Router) {
	router.HandleFunc("/", controller.Create).Method("POST")
	router.HandleFunc("/", controller.Get).Method("GET")
	router.HandleFunc("/{id}", controller.Load).Method("GET")
	router.HandleFunc("/{id}", controller.Update).Method("PUT")
	router.HandleFunc("/{id}", controller.Delete).Method("DELETE")
}