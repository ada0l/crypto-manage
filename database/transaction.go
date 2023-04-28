package database

import (
	"context"
	"crypto-manager/model"
	_ "embed"
	"github.com/jmoiron/sqlx"
	"gopkg.in/telebot.v3"
)

//go:embed sql/transaction/create.sql
var transactionCreateQuery string

func TransactionCreate(
	ctx context.Context,
	db *sqlx.DB,
	transaction model.Transaction,
) (err error) {
	_, err = db.NamedExecContext(
		ctx,
		transactionCreateQuery,
		transaction,
	)
	return err
}

func TransactionsCreate(
	ctx context.Context,
	db *sqlx.DB,
	transactions []model.Transaction,
) (err error) {
	_, err = db.NamedExecContext(
		ctx,
		transactionCreateQuery,
		transactions,
	)
	return err
}

//go:embed sql/transaction/general_info.sql
var transactionGeneralInfoQuery string

func TransactionGeneralInfo(
	ctx context.Context,
	db *sqlx.DB,
	tgUser *telebot.User,
) (info []model.TransactionGeneralInfo, err error) {
	err = db.SelectContext(
		ctx,
		&info,
		transactionGeneralInfoQuery,
		tgUser.ID,
	)
	return
}
