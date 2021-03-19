package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	//Порт запуска приложения
	port string = "8080"
	//Наша "база данных"
	db []Pizza
)

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 22,
		Price:    500.50,
		Title:    "Pepperoni",
	}

	pizza2 := Pizza{
		ID:       2,
		Diameter: 25,
		Price:    650.23,
		Title:    "BBQ",
	}
	pizza3 := Pizza{
		ID:       3,
		Diameter: 22,
		Price:    450,
		Title:    "Margaritta",
	}

	db = append(db, pizza1, pizza2, pizza3)
}

//Наша модель
type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

// Вспомогательная функция для модели (модельный метод)
func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool
	for _, p := range db {
		if p.ID == id {
			pizza = p
			found = true
			break
		}
	}
	return pizza, found
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	//Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get infos about all pizzas in database")
	writer.WriteHeader(200)            // StatusCode для запроса
	json.NewEncoder(writer).Encode(db) // Сериализация + запись в writer
}

func GetPizzaById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	//Считаем id из строки запроса и конвертируем его в int
	vars := mux.Vars(request) // {"id" : "12"}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("client trying to use invalid id param:", err)
		msg := ErrorMessage{Message: "do not use ID not supported int casting"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	log.Println("Trying to send to client pizza with id #:", id)
	pizza, ok := FindPizzaById(id)
	if ok {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(pizza)
	} else {
		msg := ErrorMessage{Message: "pizza with that id does not exists in database"}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
	}
}

func main() {
	log.Println("Trying to start REST API pizza!")
	// Инициализируем маршрутизатор
	router := mux.NewRouter()
	//1. Если на вход пришел запрос /pizzas
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	//2. Если на вход пришел запрос вида /pizza/{id}
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
}
