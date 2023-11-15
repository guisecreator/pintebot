package models

import "github.com/mymmrac/telego"

type User struct {
	Id       telego.ChatID
	Username string
	Name     string
	Phone    string
	IsAdmin  bool
}
