package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"HW1/models"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllItems(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("HW1: Get infos about all items in database")

	if len(models.DB) == 0 {
		writer.WriteHeader(403)
		json.NewEncoder(writer).Encode(models.Error{Error: "No one items found in store back"})
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.GetDBAsSlice())
}
