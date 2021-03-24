package store

import (
	"fmt"
	"log"

	"github.com/Lexa-san/spc-go2/10.JWTAuth/internal/app/models"
)

type UserRepository struct {
	store *Store
}

var (
	tableUser string = "users"
)

//Create user in database
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)
	if err := ur.store.db.QueryRow(
		query,
		u.Login,
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
		if u.Login == login {
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
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil

}

// //Create User
// func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
// 	if err := ur.store.db.QueryRow(
// 		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
// 		u.Email,
// 		u.EncryptedPassword,
// 	).Scan(&u.ID); err != nil {
// 		return nil, err
// 	} // просим вернуть id
// 	return u, nil
// }

// //Find by email
// func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
// 	return nil, nil
// }
