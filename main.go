package main

import (
	"log"
	"os"

	"github.com/Skoowshot/vecspect/telegram"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")

	b := telegram.NewTelegramBot(token, 30)
	updates, err := b.GetUpdates(-1)
	if err != nil {
		return
	}
	println(updates.LastUpdateId(-1))
}
