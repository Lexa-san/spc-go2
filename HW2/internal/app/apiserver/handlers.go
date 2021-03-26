package apiserver

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *APIServer) GetIndex(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Index API v1.0 /api/v1")
	writer.WriteHeader(http.StatusOK)
	msg := Message{
		StatusCode: http.StatusOK,
		Message:    "Hello there! It's API v1.0.",
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

func (api *APIServer) GetAllCars(writer http.ResponseWriter, req *http.Request) {
	api.logger.Info("Get All Cars GET /stock")

	initHeaders(writer)

	cars, err := api.store.Car().SelectAll()

	if err != nil {
		api.logger.Error(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing cars in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//we found nothing
	if len(cars) == 0 {
		msg := Message{
			StatusCode: 400,
			Message:    "No one autos found in DataBase",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusOK)
	//json.NewEncoder(writer).Encode(cars)
}
