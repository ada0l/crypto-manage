package bot

import (
	"context"
	"crypto-manager/database"
	"database/sql"
	"gopkg.in/telebot.v3"
)

func (h Handler) CheckExistUserMiddleware(
	next telebot.HandlerFunc,
) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		_, err := database.UserGetById(context.Background(), h.DB, c.Sender())
		if err == sql.ErrNoRows {
			return c.Send("At first you need to register using the start command")
		}
		return next(c)
	}
}
