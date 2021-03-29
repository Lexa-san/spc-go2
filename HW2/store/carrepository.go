package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Lexa-san/spc-go2/HW2/internal/app/models"
)

type CarRepository struct {
	store *Store
}

var (
	tableCar string = "car"
)

//For POST request
func (car *CarRepository) Create(c *models.Car) (*models.Car, error) {
	query := fmt.Sprintf("INSERT INTO %s (mark, max_speed, distance, stock, handler) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING car_id", tableCar)
	if err := car.store.db.QueryRow(query, c.Mark, c.MaxSpeed, c.Distance, c.Stock, c.Handler).Scan(&c.ID); err != nil {
		return nil, err
	}
	return c, nil
}

//Helper fo Find by id and GET by id request
//func (car *CarRepository) FindCarById(id int) (*models.Car, bool, error) {
//	cCars, err := car.SelectAll()
//	founded := false
//	if err != nil {
//		return nil, founded, err
//	}
//	var cCarFinded *models.Car
//	for _, a := range cCars {
//		if a.ID == id {
//			cCarFinded = a
//			founded = true
//		}
//	}
//
//	return cCarFinded, founded, nil
//
//}

//Get all request and helper for FindByID
func (car *CarRepository) SelectAll() ([]*models.Car, error) {
	query := fmt.Sprintf("SELECT c.car_id, c.mark, max_speed, c.distance, c.handler, c.stock FROM %s as c", tableCar)
	rows, err := car.store.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	cCars := make([]*models.Car, 0)
	for rows.Next() {
		a := models.Car{}
		err := rows.Scan(&a.ID, &a.Mark, &a.MaxSpeed, &a.Distance, &a.Handler, &a.Stock)
		if err != nil {
			log.Println(err)
			continue
		}
		cCars = append(cCars, &a)
	}
	return cCars, nil
}

func (car *CarRepository) SelectOneByMark(mark string) (*models.Car, bool, error) {
	a := models.Car{}
	query := fmt.Sprintf("SELECT "+
		"c.car_id, "+
		"c.mark,"+
		"c.max_speed, "+
		"c.distance, "+
		"c.handler, "+
		"c.stock "+
		"FROM %s as c "+
		"WHERE c.mark ilike $1 ", tableCar)

	err := car.store.db.QueryRow(query, mark).Scan(&a.ID, &a.Mark, &a.MaxSpeed, &a.Distance, &a.Handler, &a.Stock)
	switch {
	case err == sql.ErrNoRows:
		return &a, false, nil
	case err != nil:
		return &a, false, err
	}

	return &a, true, nil
}

//UpdateByID Car
func (car *CarRepository) UpdateByID(c *models.Car, id int) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET "+
		"mark = $2, "+
		"max_speed = $3, "+
		"distance = $4, "+
		"stock = $5, "+
		"handler = $6 "+
		"WHERE car_id = $1", tableCar)
	res, err := car.store.db.Exec(query, id, c.Mark, c.MaxSpeed, c.Distance, c.Stock, c.Handler)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	return true, nil
}

//DeleteByID Car
func (car *CarRepository) DeleteById(id int) (bool, error) {
	query := fmt.Sprintf("DELETE from %s WHERE car_id = $1", tableCar)
	res, err := car.store.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	return true, nil
}
