package bot

import (
	"context"
	"crypto-manager/database"
	"crypto-manager/model"
	"fmt"
	"gopkg.in/telebot.v3"
)

func (h Handler) onAdd(c telebot.Context) error {
	args := c.Args()
	if len(args) != 4 {
		return c.Send(
			fmt.Sprintf(
				"4 arguments were expected, but received %d",
				len(args),
			),
		)
	}
	transaction, err := model.NewTransaction(c.Sender(), args[0], args[1], args[2], args[3])
	if err != nil {
		return c.Send(fmt.Sprint(err))
	}
	ctx := context.Background()
	err = database.TransactionCreate(ctx, h.DB, transaction)
	if err != nil {
		return c.Send(fmt.Sprintf("Error during create transaction. %v", err))
	}
	return c.Send("Success")
}
