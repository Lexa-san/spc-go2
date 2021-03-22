package utils

import (
	"HW1/handlers"
	"github.com/gorilla/mux"
)

func BuildItemResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetItemById).Methods("GET")
	router.HandleFunc(prefix+"/{id}", handlers.CreateItem).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateItemById).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteItemById).Methods("DELETE")
}

func BuildManyItemResourcePrefix(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, handlers.GetAllItems).Methods("GET")
}
