package models

var (
	DB     []Item
	LastId int
)

type Item struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price"`
}

func init() {
	item := Item{
		ID:     GetNextId(),
		Title:  "Тапочки",
		Amount: 100,
		Price:  19.99,
	}
	DB = append(DB, item)

}

func GetNextId() int {
	LastId++
	return LastId
}

func FindItemById(id int) (Item, bool) {
	var item Item
	var found bool
	for _, b := range DB {
		if b.ID == id {
			item = b
			found = true
			break
		}
	}
	return item, found
}

func DelItemById(id int) bool {
	var found bool
	for i, b := range DB {
		if b.ID == id {
			found = true
			DB[i] = DB[len(DB)-1]
			DB = DB[:len(DB)-1]
			break
		}
	}
	return found
}

func UpdateItemById(id int, item Item) bool {
	var found bool
	for i, b := range DB {
		if b.ID == id {
			found = true
			item.ID = b.ID
			DB[i] = item
			break
		}
	}
	return found
}
