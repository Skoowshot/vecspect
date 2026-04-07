package logic

import (
	"github.com/Skoowshot/vecspect/domain"
	"github.com/Skoowshot/vecspect/telegram"
)

type App struct {
	bot      *telegram.TelegramBot
	listener *telegram.DefaultUpdateListener
}

func NewApp(token string) *App {
	a := &App{
		bot: telegram.NewTelegramBot(token, 30),
	}
	a.listener = telegram.NewDefaultUpdateListener(
		a.OnMessage,
	)
	return a
}

func (a *App) Start() {
	a.bot.PollUpdates(a.listener)
}

func (a *App) OnMessage(message *domain.Message) {
	println("received message from " + message.FromUser.Username + ": " + message.Text)

	if replyTo := message.ReplyTo; replyTo != nil {
		a.OnReply(message, replyTo)
	}
}

func (a *App) OnReply(message *domain.Message, replyTo *domain.Message) {
	println("> reply to @" + replyTo.FromUser.Username + ": " + replyTo.Text)
}
