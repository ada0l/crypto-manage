package bot

import (
	"github.com/jmoiron/sqlx"
	"time"

	"gopkg.in/telebot.v3"
)

func New(token string, db *sqlx.DB) (*telebot.Bot, error) {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		return nil, err
	}
	handler := NewHandler(b, db)
	b.Handle("/start", handler.onStart)

	checkExistUser := b.Group()
	checkExistUser.Use(handler.CheckExistUserMiddleware)
	checkExistUser.Handle("/add", handler.onAdd)
	checkExistUser.Handle("/export", handler.onExport)
	checkExistUser.Handle(telebot.OnDocument, handler.onImport)
	checkExistUser.Handle("/general_info", handler.onGeneralInfo)

	return b, nil
}
