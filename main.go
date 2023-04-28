package main

import (
	"crypto-manager/bot"
	"crypto-manager/database"
	"fmt"
	"os"
)

func main() {
	db, err := database.GetDB(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(fmt.Sprintf("Failed connect to database: %v", err))
	}
	b, err := bot.New(os.Getenv("TELEGRAM_BOT"), db)
	if err != nil {
		panic(fmt.Sprintf("Failed to create telegram bot: %v", err))
	}
	b.Start()
}
