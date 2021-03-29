package apiserver

import (
	"github.com/Lexa-san/spc-go2/HW2/internal/app/middleware"
	"github.com/Lexa-san/spc-go2/HW2/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	prefix string = "/api/v1"
)

// type for APIServer object for instancing server
type APIServer struct {
	//Unexported field
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//APIServer constructor
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server and connection to db and logger confs
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting api server at port: ", s.config.BindAddr)
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

//func for configureate logger, should be unexported
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return nil
	}
	s.logger.SetLevel(level)

	return nil
}

//func for configure Router
func (s *APIServer) configureRouter() {
	//test method to check API version
	s.router.HandleFunc(prefix, s.GetIndex).Methods("GET")

	//user methods
	s.router.HandleFunc(prefix+"/register", s.PostUserRegister).Methods("POST")
	s.router.HandleFunc(prefix+"/auth", s.PostToAuth).Methods("POST")

	//method with auth
	s.router.Handle(prefix+"/stock", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.GetAllCars),
	)).Methods("GET")
	s.router.Handle(prefix+"/auto/{mark}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.GetCarByMark),
	)).Methods("GET")
	s.router.Handle(prefix+"/auto/{mark}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.CreateCar),
	)).Methods("POST")
	s.router.Handle(prefix+"/auto/{mark}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.UpdateCar),
	)).Methods("PUT")
	s.router.Handle(prefix+"/auto/{mark}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(s.DeleteCar),
	)).Methods("DELETE")

	////
	//s.router.HandleFunc(prefix+"/articles"+"/{id}", s.DeleteArticleById).Methods("DELETE")
	//s.router.HandleFunc(prefix+"/articles", s.PostArticle).Methods("POST")
}

//configureStore method
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
