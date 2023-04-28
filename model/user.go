package model

import "gopkg.in/telebot.v3"

type User struct {
	Id int64 `db:"id"`
}

func NewUser(tgUser *telebot.User) User {
	return User{Id: tgUser.ID}
}
