package main

import (
	"log"
	"os"

	"github.com/Skoowshot/vecspect/domain"
	"github.com/Skoowshot/vecspect/logic"
	"github.com/joho/godotenv"
)

type Listener struct{}

func (*Listener) OnUpdate(update *domain.Update) {
	println(update.Message.Text)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	
	a := logic.NewApp(token)
	a.Start()
}
