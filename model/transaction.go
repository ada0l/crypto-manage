package model

import (
	"errors"
	"gopkg.in/telebot.v3"
	"strconv"
	"time"
)

type (
	Transaction struct {
		Id           int64     `db:"id"`
		UserId       int64     `db:"telegram_user_id"`
		PurchaseDate time.Time `db:"purchase_date"`
		Asset        string    `db:"asset"`
		Price        float64   `db:"price"`
		Amount       float64   `db:"amount"`
	}

	TransactionGeneralInfo struct {
		Asset    string  `db:"asset"`
		Amount   float64 `db:"amount"`
		AvgPrice float64 `db:"avg_price"`
		Spend    float64 `db:"spend"`
		Percent  float64 `db:"percent"`
	}
)

func NewTransaction(
	tgUser *telebot.User,
	purchaseData_,
	asset,
	price_,
	amount_ string,
) (transaction Transaction, err error) {
	purchaseDate, err := time.Parse("2006-01-02", purchaseData_)
	if err != nil {
		return transaction, errors.New("the wrong date format")
	}
	price, err := strconv.ParseFloat(price_, 64)
	if err != nil {
		return transaction, errors.New("the wrong price format")
	}
	amount, err := strconv.ParseFloat(amount_, 64)
	if err != nil {
		return transaction, errors.New("the wrong amount format")
	}
	transaction = Transaction{
		UserId:       tgUser.ID,
		Asset:        asset,
		PurchaseDate: purchaseDate,
		Price:        price,
		Amount:       amount,
	}
	return
}
