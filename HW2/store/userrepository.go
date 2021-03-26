package store

import (
	"fmt"
	"log"

	"github.com/Lexa-san/spc-go2/HW2/internal/app/models"
)

type UserRepository struct {
	store *Store
}

var (
	tableUser string = "user"
)

//Create user in database
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password) VALUES ($1, $2) RETURNING id", tableUser)
	if err := ur.store.db.QueryRow(
		query,
		u.Username,
		u.Password,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//Find by login
func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var userFinded *models.User
	for _, u := range users {
		if u.Username == login {
			userFinded = u
			founded = true
			break
		}
	}
	return userFinded, founded, nil
}

//Select All
func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Username, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil

}
