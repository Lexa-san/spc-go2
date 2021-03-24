package models

//User model defeniton
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
