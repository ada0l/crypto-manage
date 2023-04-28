package bot

import (
	"gopkg.in/telebot.v3"
)

func (h Handler) onExport(c telebot.Context) error {
	return c.Send("Success")
}
