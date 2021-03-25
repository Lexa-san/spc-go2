package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

//Пытаемся откунфигурировать наш API инстанс (а конкретнее - поле logger)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

//Пытаемся отконфигурировать маршрутизатор (а конкретнее поле router API)
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is API of 9.TaskSW!"))
	})
	a.router.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is API v1.0 of 9.TaskSW!"))
	})

	//a.router.HandleFunc(prefix+"/q", a.GetQuadratic).Methods("GET")
	a.router.HandleFunc(prefix+"/grab", a.GrabQuadratic).Methods("POST")
	a.router.HandleFunc(prefix+"/solve", a.SolveQuadratic).Methods("GET")
}
