package bot

import (
	"bytes"
	"context"
	"crypto-manager/database"
	"crypto-manager/log"
	"text/template"

	"gopkg.in/telebot.v3"
)

func (h Handler) onGeneralInfo(c telebot.Context) error {
	ctx := context.Background()
	report, err := database.TransactionGeneralInfo(ctx, h.DB, c.Sender())
	if err != nil {
		return c.Send("Error")
	}
	t := template.New("generalInfo.tmpl")
	t, err = t.ParseFiles("template/generalInfo.tmpl")
	if err != nil {
		log.Error.Println(err)
		return c.Send("Error during load template file")
	}
	var answer bytes.Buffer
	if err = t.Execute(&answer, report); err != nil {
		log.Error.Println(err)
		return c.Send("Error during execute template file")
	}
	return c.Send(answer.String())
}
