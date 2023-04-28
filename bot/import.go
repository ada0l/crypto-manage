package bot

import (
	"context"
	"crypto-manager/database"
	"crypto-manager/model"
	"encoding/csv"
	"fmt"

	"gopkg.in/telebot.v3"
)

func (h Handler) onImport(c telebot.Context) error {
	document := c.Message().Document
	reader, err := c.Bot().File(&document.File)
	if err != nil {
		return c.Send("Error during file open")
	}
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return c.Send("Error during file processing")
	}
	transactions := make([]model.Transaction, 0)
	for i, record := range records {
		if len(record) != 4 {
			return c.Send(
				"4 arguments were expected, but received %d in %d row",
				len(record),
				i,
			)
		}
		transaction, err := model.NewTransaction(c.Sender(), record[0], record[1], record[2], record[3])
		if err != nil {
			return c.Send(fmt.Sprintf("%v in %d row", err, i))
		}
		transactions = append(transactions, transaction)
	}
	ctx := context.Background()
	err = database.TransactionsCreate(ctx, h.DB, transactions)
	if err != nil {
		return c.Send(fmt.Sprint(err))
	}
	return c.Send("Success")
}
