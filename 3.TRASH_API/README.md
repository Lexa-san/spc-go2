## Лекция 3. Простейший API и термины

**Задача**: создать простой REST API , который будет позволять получать информацию про пиццу.

## Шаг 1. Желаемый функционал.
Хотим собрать простей веб-сервер, который будет взаимодействовать с окружающим миром через API поддерживающий REST.

### Шаг 1.1 Идея
Хочется, чтобы наш сервер помогал клиентам узнавать следующую информацию:
* Какая пицца есть в наличии?
* Информация, про какую-то конкретную пиццу.

### Шаг 1.2 Виды запросов, поддерживаемые api
Будет существовать и поддерживаться 2 запроса:
* ```http://localhost:8080/pizzas``` - возвращает json со всеми пиццами в наличии
* ```http://localhost:8080/pizza/{id}``` - возвращает информацию про пиццу с ```id``` в случае если она имеется в наличии, или сообщает клиенту, что такой пиццы нет.


### Шаг 2. Реализация
* Для начала создадим файл ```main.go```
* Сразу инициализируем ```go mod init 3.Trash_API```

### Шаг 2.2. Базовый скелет
```
package main

import (
	"log"
	"net/http"
)

var (
	port string = "8080"
)

func main() {
	log.Println("Trying to start REST API pizza!")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

```
**ВАЖНО**: для остановки приложения используем ```Ctrl+C``` (остановка текущего процесса и ОСВОБОЖДЕНИЕ РЕСУРСОВ).

### Шаг 2.3 Маршрутизатор и исполнители
***Маршрутизатор (router)*** - это экземпляр, который имеет внутренний функционал , заключающийся в следующем:
* принимает на вход адрес запроса (по сути это строка ```http://localhost:8080/pizzas```) и вызывает исполнителя, который будет ассоциирован с этим запросом.

***Исполнитель (handler)*** - это функция/метод, котоырй вызывает маршрутизатором.

Для того, чтобы удобно работать с маршрутизатором и не писать его с нуля. Для этого установим уже готовую библиотеку:  ```github.com/gorilla/mux``` . А устаналвивается запросом ```go get -u github.com/gorilla/mux```

```
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
)

func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {}
func GetPizzaById(writer http.ResponseWriter, request *http.Request) {}

func main() {
	log.Println("Trying to start REST API pizza!")
	// Инициализируем маршрутизатор
	router := mux.NewRouter()
	//1. Если на вход пришел запрос /pizzas
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	//2. Если на вход пришел запрос вида /pizza/{id}
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

```

На данный момент нами написан базовый скелет функционала API (сейчас отсутствует хранилище данных и внутренняя логика), но тем не менее, сервер конфигурируется и уже запускается.

### Шаг 2.4 Создаем модель данных
В качестве хранилища выберем слайс с экземплярами пиццы.
Для этого реализуем следующий функционал:
```
var db []Pizza


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
```

Определили базу данных ```db```. Определили структуру ```Pizza``` с сопоставлением полей, а также определили функцию, которая будет просматривать слайс и говорить, есть ли в нем нужная нам пицца, или ее нет.

### Шаг 2.5 Реализуем исполнителей (handlers)
Поскольку мы собираемся запускать наш сервер , как поддерживающий REST API архитектуру, нужно, чтобы в теле каждого ответа фигурировала информация про то, каким образом наш сервер общается!
```
func GetAllPizzas(writer http.ResponseWriter, request *http.Request) {
	//Прописывать хедеры .
	writer.Header().Set("Content-Type", "application/json")
	log.Println("Get infos about all pizzas in database")
	writer.WriteHeader(200)            // StatusCode для запроса
	json.NewEncoder(writer).Encode(db) // Сериализация + запись в writer
}

```

Данный обработчик возвращает сериализованный json полученный на основе ```db``` (слайса Пицц)
```
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

```