package models

var (
	DB     []Book
	LastId int
)

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		ID:            GetNextId(),
		Title:         "Lord of the Rings. Vol.1",
		YearPublished: 1978,
		Author: Author{
			Name:     "J.R",
			LastName: "Tolkin",
			BornYear: 1892,
		},
	}
	DB = append(DB, book1)

}

func GetNextId() int {
	LastId++
	return LastId
}

func FindBookById(id int) (Book, bool) {
	var book Book
	var found bool
	for _, b := range DB {
		if b.ID == id {
			book = b
			found = true
			break
		}
	}
	return book, found
}

func DelBookById(id int) bool {
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

func UpdateBookById(id int, book Book) bool {
	var found bool
	for i, b := range DB {
		if b.ID == id {
			found = true
			book.ID = b.ID
			DB[i] = book
			break
		}
	}
	return found
}
