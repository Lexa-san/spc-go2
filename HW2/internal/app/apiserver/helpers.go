package apiserver

import (
	"encoding/json"
	"net/http"
)

func alertBadJSON(api *APIServer, w http.ResponseWriter, err error) {
	api.logger.Info("Invalid json recieved from client: ", err)
	msg := Message{
		StatusCode: 400,
		Message:    "Provided json is invalid",
		IsError:    true,
	}
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(msg)
}

func alertDBError(api *APIServer, writer http.ResponseWriter, err error) {
	api.logger.Info("Troubles while accessing database table (user) with Username. err: ", err)
	msg := Message{
		StatusCode: 500,
		Message:    "We have some troubles to accessing database. Try again",
		IsError:    true,
	}
	writer.WriteHeader(500)
	json.NewEncoder(writer).Encode(msg)
}

func alertUserDoesNotExist(api *APIServer, writer http.ResponseWriter) {
	api.logger.Info("User with that login does not exists")
	msg := Message{
		StatusCode: 400,
		Message:    "User with that login does not exists in database. Try register first",
		IsError:    true,
	}
	writer.WriteHeader(400)
	json.NewEncoder(writer).Encode(msg)
}

func alertBadAuthCredentials(api *APIServer, writer http.ResponseWriter) {
	api.logger.Info("Invalid credentials to auth")
	msg := Message{
		StatusCode: 404,
		Message:    "Your password is invalid",
		IsError:    true,
	}
	writer.WriteHeader(404)
	json.NewEncoder(writer).Encode(msg)
}

func alertTokenGenerate(api *APIServer, writer http.ResponseWriter, err error) {
	api.logger.Error("Can not claim jwt-token. Err: ", err)
	msg := Message{
		StatusCode: 500,
		Message:    "We have some troubles. Try again",
		IsError:    true,
	}
	writer.WriteHeader(500)
	json.NewEncoder(writer).Encode(msg)
}

func alertCarDoesNotExist(api *APIServer, writer http.ResponseWriter) {
	api.logger.Info("Can not find car with that ID in database")
	msg := Message{
		StatusCode: 404,
		Message:    "Car with that mark not found",
		IsError:    true,
	}

	writer.WriteHeader(404)
	json.NewEncoder(writer).Encode(msg)
}

func alertCarAlreadyExist(api *APIServer, writer http.ResponseWriter) {
	api.logger.Info("Car with that mark exists in database")
	msg := Message{
		StatusCode: 400,
		Message:    "Car with that mark exists",
		IsError:    true,
	}

	writer.WriteHeader(400)
	json.NewEncoder(writer).Encode(msg)
}
