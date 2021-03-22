package handlers

import (
	"encoding/json"
	"fmt"
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
		msg := models.Error{Error: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book, ok := models.FindItemById(id)
	log.Println("HW1: Get item with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Error{Error: "Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateItem(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("HW1: Creating new item ....")
	var (
		item models.Item
		id   int
		err  error
	)

	id, err = strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("HW1: error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Error{Error: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&item)
	if err != nil {
		msg := models.Error{Error: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if err = models.AddItem(item, id); err != nil {
		log.Println("HW1: error with store new Item:", err)
		writer.WriteHeader(400)
		msg := models.Error{Error: fmt.Sprint(err)}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(201)
	msg := models.Message{Message: "Item created"}
	json.NewEncoder(writer).Encode(msg)
}

func UpdateItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("HW1: Updating item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("HW1: error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Error{Error: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok := models.FindItemById(id)
	var newItem models.Item
	if !ok {
		log.Println("HW1: item not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.Error{Error: "item with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&newItem)
	if err != nil {
		msg := models.Error{Error: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if newItem, err = models.UpdateItemById(id, newItem); err != nil {
		log.Println("HW1: item was not updated. id:", err)
		writer.WriteHeader(404)
		msg := models.Error{Error: fmt.Sprint(err)}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	json.NewEncoder(writer).Encode(newItem)
}

func DeleteItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("HW1: Deleting item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("HW1: error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Error{Error: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok := models.FindItemById(id)
	if !ok {
		log.Println("HW1: item was not found. id:", id)
		writer.WriteHeader(404)
		msg := models.Error{Error: "Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if err = models.DelItemById(id); err != nil {
		log.Println("HW1: item was not deleted. id:", id)
		writer.WriteHeader(500)
		msg := models.Error{Error: fmt.Sprint(err)}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := models.Message{Message: "Item deleted"}
	json.NewEncoder(writer).Encode(msg)
}
