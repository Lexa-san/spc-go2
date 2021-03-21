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
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.GetDBAsSlice())
}
