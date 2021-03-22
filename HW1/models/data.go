package models

import (
	"errors"
	"log"
)

var (
	DB     map[int]Item
	LastId int
)

type Item struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price"`
}

func init() {
	DB = make(map[int]Item)
	log.Println("HW1: init DB:", DB)
	id := GenerateNextId()
	item := Item{
		ID:     id,
		Title:  "Тапочки",
		Amount: 100,
		Price:  19.99,
	}
	DB[id] = item
	log.Println("HW1: add first item to DB:", DB)
}

// Generate nex ID that can be used to new Item in DB
func GenerateNextId() int {
	LastId++
	return LastId
}

// Find Item in DB by ID
func FindItemById(id int) (Item, bool) {
	item, found := DB[id]
	return item, found
}

// Add a new Item to DB and return its ID
func AddItem(item Item, id int) error {
	if _, ok := DB[id]; ok {
		return errors.New("Item with that id already exists")
	}

	item.ID = id
	DB[id] = item
	return nil
}

// Delete Item from DB by ID
func DelItemById(id int) error {
	if _, ok := DB[id]; !ok {
		return errors.New("Item with that id not found")
	}
	delete(DB, id)
	return nil
}

// update Item in DB by ID
func UpdateItemById(id int, item Item) (Item, error) {
	if _, ok := DB[id]; !ok {
		return Item{}, errors.New("Item with that id not found")
	}

	item.ID = id
	DB[id] = item
	return DB[id], nil
}

// convert DB from map to slice to be able to create JSON
func GetDBAsSlice() []Item {
	result := make([]Item, 0, len(DB))
	for _, item := range DB {
		result = append(result, item)
	}
	return result
}
