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
