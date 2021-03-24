package models

//Article model defenition
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
