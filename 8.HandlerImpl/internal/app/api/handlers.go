package api

import (
	"encoding/json"
	"net/http"

	"github.com/Lexa-san/spc-go2/8.HandlersImpl/internal/app/models"
)

//Вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

//Возвращает все статьи из бд на данный момент
func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	//Инициализируем хедеры
	initHeaders(writer)
	//Логируем момент начало обработки запроса
	api.logger.Info("Get All Artiles GET /api/v1/articles")
	//Пытаемся что-то получить от бд
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		//Что делаем, если была ошибка на этапе подключения?
		api.logger.Info("Error while Articles.SelectAll : ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Troubles while creating new article:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again.",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {}

func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {}

func (api *API) PostUserRegister(writer http.ResponseWriter, req *http.Request) {}
