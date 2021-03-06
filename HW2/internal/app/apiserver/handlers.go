package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/Lexa-san/spc-go2/HW2/internal/app/middleware"
	"github.com/Lexa-san/spc-go2/HW2/internal/app/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

//Init common HTTP headers
func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

//Get index router of API v1.0
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

//Register user
func (api *APIServer) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post User Register POST /api/v1/register")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		alertBadJSON(api, writer, err)
		return
	}

	//Пытаемся найти пользователя с таким логином в бд
	_, ok, err := api.store.User().SelectOneByLogin(user.Username)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}

	//Смотрим, если такой пользователь уже есть - то никакой регистрации мы не делаем!
	if ok {
		api.logger.Info("User with that Username already exists")
		msg := Message{
			StatusCode: 400,
			Message:    "User already exists",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//Теперь пытаемся добавить в бд
	_, err = api.store.User().Create(&user)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}

	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User created. Try to auth"),
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}

func (api *APIServer) PostToAuth(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post to Auth POST /api/v1/auth")
	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	//id bad json
	if err != nil {
		alertBadJSON(api, writer, err)
		return
	}

	userInDB, ok, err := api.store.User().SelectOneByLogin(userFromJson.Username)
	//if trouble with db
	if err != nil {
		alertDBError(api, writer, err)
		return
	}

	//if user does not exist
	if !ok {
		alertUserDoesNotExist(api, writer)
		return
	}
	//check auth credentials
	if userInDB.Password != userFromJson.Password {
		alertBadAuthCredentials(api, writer)
		return
	}

	//generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["admin"] = true
	claims["name"] = userInDB.Username
	tokenString, err := token.SignedString(middleware.SecretKey)
	//if trouble with token generation
	if err != nil {
		alertTokenGenerate(api, writer, err)
		return
	}

	//if OK
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)
}

//Get all cars
func (api *APIServer) GetAllCars(writer http.ResponseWriter, req *http.Request) {
	api.logger.Info("Get All Cars GET /api/v1/stock")

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
	json.NewEncoder(writer).Encode(cars)
}

func (api *APIServer) GetCarByMark(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Car by mark /api/v1/auto/{mark}")
	mark := mux.Vars(req)["mark"]
	car, ok, err := api.store.Car().SelectOneByMark(mark)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}
	if !ok {
		alertCarDoesNotExist(api, writer)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(car)

}

func (api *APIServer) CreateCar(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Create Car POST /auto/{mark}")
	var err error

	mark := mux.Vars(req)["mark"]
	_, ok, err := api.store.Car().SelectOneByMark(mark)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}
	if ok {
		alertCarAlreadyExist(api, writer)
		return
	}

	var car models.Car
	err = json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		alertBadJSON(api, writer, err)
		return
	}

	car.Mark = mark
	a, err := api.store.Car().Create(&car)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)

}

func (api *APIServer) UpdateCar(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("UpdateByID Car PUT /auto/{mark}")
	var err error
	var ok bool

	mark := mux.Vars(req)["mark"]
	foundCar, ok, err := api.store.Car().SelectOneByMark(mark)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}
	if !ok {
		alertCarDoesNotExist(api, writer)
		return
	}

	var car models.Car
	err = json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		alertBadJSON(api, writer, err)
		return
	}

	ok, err = api.store.Car().UpdateByID(&car, foundCar.ID)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}
	if !ok {
		alertCarDoesNotExist(api, writer)
		return
	}
	msg := Message{
		IsError:    false,
		Message:    "Auto updated",
		StatusCode: 202,
	}
	writer.WriteHeader(202)
	json.NewEncoder(writer).Encode(msg)

}

func (api *APIServer) DeleteCar(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("DeleteByID Car DELETE /auto/{mark}")
	var err error
	var ok bool

	mark := mux.Vars(req)["mark"]
	foundCar, ok, err := api.store.Car().SelectOneByMark(mark)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}
	if !ok {
		alertCarDoesNotExist(api, writer)
		return
	}

	ok, err = api.store.Car().DeleteById(foundCar.ID)
	if err != nil {
		alertDBError(api, writer, err)
		return
	}
	if !ok {
		alertCarDoesNotExist(api, writer)
		return
	}
	msg := Message{
		IsError:    false,
		Message:    "Auto deleted",
		StatusCode: 202,
	}
	writer.WriteHeader(202)
	json.NewEncoder(writer).Encode(msg)

}
