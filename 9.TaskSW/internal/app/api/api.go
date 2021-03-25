package api

import (
	"github.com/Lexa-san/spc-go2/9.TaskSW/internal/app/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	prefix string = "/api/v1"
	//DB map[string]int
	DB *models.Quadratic
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

func NewDB() *models.Quadratic {
	return new(models.Quadratic)
}

// Start http server/configure loggers, router, database connection and etc....
func (api *API) Start() error {

	//Trying to confugre logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	//Подтверждение того, что логгер сконфигурирован
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	//DB = make(map[string]int)
	DB = NewDB()
	api.logger.Info("DB pointer:", DB, "DB value:", *DB)

	//Конфигурируем маршрутизатор
	api.configureRouterField()
	//На этапе валидного завршениея стратуем http-сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
