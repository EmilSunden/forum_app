package models

type Auth struct {
	Username string `json:"username"`
	Password string `json:"-password"`
}
