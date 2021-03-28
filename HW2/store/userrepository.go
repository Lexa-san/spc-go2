package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Lexa-san/spc-go2/HW2/internal/app/models"
)

type UserRepository struct {
	store *Store
}

var (
	tableUser string = "users"
)

//Create user in database
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING user_id", tableUser)
	if err := ur.store.db.QueryRow(
		query,
		u.Username,
		u.Password,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//Select one User from DB
func (ur *UserRepository) SelectOneByLogin(login string) (*models.User, bool, error) {
	a := models.User{}
	query := fmt.Sprintf("SELECT "+
		"c.user_id, "+
		"c.login, "+
		"c.password "+
		"FROM %s as c "+
		"where c.login ilike $1", tableUser)
	log.Println(query)

	err := ur.store.db.QueryRow(query, login).Scan(&a.ID, &a.Username, &a.Password)
	switch {
	case err == sql.ErrNoRows:
		return &a, false, nil
	case err != nil:
		return &a, false, err
	}
	return &a, true, nil
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

//Find by login
//func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
//	users, err := ur.SelectAll()
//	var founded bool
//	if err != nil {
//		return nil, founded, err
//	}
//	var userFinded *models.User
//	for _, u := range users {
//		if u.Username == login {
//			userFinded = u
//			founded = true
//			break
//		}
//	}
//	return userFinded, founded, nil
//}
