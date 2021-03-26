package models

//Car models...
type Car struct {
	ID       int    `json:"id"`
	Mark     string `json:"mark"`
	MaxSpeed int    `json:"max_speed"`
	Distance int    `json:"distance"`
	Handler  string `json:"content"`
	Stock    string `json:"content"`
}
