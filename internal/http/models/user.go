package models

type User struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"password"`
}
