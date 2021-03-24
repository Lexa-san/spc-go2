package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	prefix string = "/api/v1"
	//qStore *models.Quadratic
	DB map[string]int
)

//Base API server instance description
type API struct {
	//!!! UNEXPORTED FIELD !!!
	config *Congig
	logger *logrus.Logger
	router *mux.Router
}

//API constructor: build base API instance
func New(config *Congig) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server/configure loggers, router, database connection and etc....
func (api *API) Start() error {

	//Trying to confugre logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	//Подтверждение того, что логгер сконфигурирован
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	DB = make(map[string]int)

	//Конфигурируем маршрутизатор
	api.configureRouterField()
	//На этапе валидного завршениея стратуем http-сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
