package bot

import (
	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v3"
)

type Handler struct {
	Bot *telebot.Bot
	DB  *sqlx.DB
}

func NewHandler(bot *telebot.Bot, db *sqlx.DB) *Handler {
	return &Handler{
		Bot: bot,
		DB:  db,
	}
}
