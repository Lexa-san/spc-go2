package api

import (
	"github.com/Lexa-san/spc-go2/7.ServerAndDB2/storage"
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
		w.Write([]byte("Hello! This is rest api!"))
	})
}

//Пытаемся отконфигурировать наше хранилище (storage API)
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединениение, если невозможно - возвращаем ошибку!
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
