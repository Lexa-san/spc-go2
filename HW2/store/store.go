package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Instance of store
type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
	carRepository  *CarRepository
}

// Constructor for store
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open store method
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	//Проверим, что все ок. Реально соединение тут не создается. Соединение только при первом вызове
	//db.Ping() // Пустой SELECT *
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("Connection to db successfully")

	var str string
	if err = db.QueryRow("select current_time").Scan(&str); err != nil {
		log.Fatal("select error: ", err)
	}
	log.Println("select success: ", str)

	return nil
}

//Close store method
func (s *Store) Close() {
	s.db.Close()
}

//Public for UserRepo
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

//Public for CarRepo
func (s *Store) Car() *CarRepository {
	if s.carRepository != nil {
		return s.carRepository
	}
	s.carRepository = &CarRepository{
		store: s,
	}
	return s.carRepository
}
