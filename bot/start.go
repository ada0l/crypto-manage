package bot

import (
	"context"
	"crypto-manager/database"
	"crypto-manager/model"
	"gopkg.in/telebot.v3"
)

func (h Handler) onStart(c telebot.Context) error {
	ctx := context.Background()
	_, err := database.UserGetById(ctx, h.DB, c.Sender())
	if err == nil {
		return c.Send("You already registered")
	}
	err = database.UserCreate(ctx, h.DB, model.NewUser(c.Sender()))
	if err != nil {
		return c.Send("Error during create user")
	}
	return c.Send("Welcome")
}
