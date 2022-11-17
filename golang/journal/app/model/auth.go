package model

type Auth struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
}
