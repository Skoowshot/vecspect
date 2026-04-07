package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Skoowshot/vecspect/logic"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")

	a := logic.NewApp(token)
	go a.Start()

	println("bot started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	println("shutting down")
}
