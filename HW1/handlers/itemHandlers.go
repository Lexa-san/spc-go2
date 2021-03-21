package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"HW1/models"
	"github.com/gorilla/mux"
)

func GetItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("HW1: error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book, ok := models.FindItemById(id)
	log.Println("HW1: Get item with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "item with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateItem(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("HW1: Creating new item ....")
	var item models.Item

	err := json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	item.ID = models.GetNextId()
	models.DB = append(models.DB, item)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(item)
}

func UpdateItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("HW1: Updating item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("HW1: error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok := models.FindItemById(id)
	var newItem models.Item
	if !ok {
		log.Println("HW1: item not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "item with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&newItem)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	ok = models.UpdateItemById(id, newItem)
	if !ok {
		log.Println("HW1: item was not updated. id:", id)
		writer.WriteHeader(204)
		//msg := models.Message{Message: "item with that ID was not updated in database"}
		//json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	msg := models.Message{Message: "successfully updated requested item"}
	json.NewEncoder(writer).Encode(msg)
}

func DeleteItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("HW1: Deleting item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("HW1: error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok := models.FindItemById(id)
	if !ok {
		log.Println("HW1: item not found in database. id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "item with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	ok = models.DelItemById(id)
	if !ok {
		log.Println("HW1: item not found in data base . id :", id)
		writer.WriteHeader(204)
		//msg := models.Message{Message: "item with that ID does not deleted in database"}
		//json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	msg := models.Message{Message: "successfully deleted requested item"}
	json.NewEncoder(writer).Encode(msg)
}
