package logic

import (
	"log"

	"github.com/Skoowshot/vecspect/domain"
	"github.com/Skoowshot/vecspect/logic/train"
	"github.com/Skoowshot/vecspect/telegram"
)

const pollTimeout = 30
const workerCount = 4
const queueLength = 100

type App struct {
	bot      *telegram.TelegramBot
	listener *telegram.DefaultUpdateListener
	learn    *train.TrainingOrchestrator
}

func NewApp(token string) *App {
	a := &App{
		bot:   telegram.NewTelegramBot(token, pollTimeout),
		learn: train.NewOrchestrator(workerCount, queueLength),
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
	if message.Text == "" {
		log.Println("received empty message, skipping")
		return
	}

	log.Printf("received message from %v: %v", message.FromUser.Username, message.Text)
	log.Printf("| chat %v, forum: %v, thread: %v", message.Chat.ChatId, message.Chat.IsForum, message.MessageThreadId)

	if replyTo := message.ReplyTo; replyTo != nil {
		if replyTo.Text != "" {
			a.OnReply(message, replyTo)
		}
	} else {
		a.OnRegularMessage(message)
	}
}

func (a *App) OnRegularMessage(message *domain.Message) {
	chatId := message.Chat.ChatId
	originalText := message.ReplyTo.Text

	msg := domain.NewTrainMessage(chatId, originalText)
	a.learn.Push(msg)
}

func (a *App) OnReply(message *domain.Message, replyTo *domain.Message) {
	log.Printf("| reply to %v: %v", replyTo.FromUser.Username, replyTo.Text)
	if message.FromUser.IsBot || replyTo.FromUser.IsBot {
		log.Println("| skipping, one of users (reply or from) is a bot")
		return
	}

	chatId := message.Chat.ChatId
	replyText := message.Text
	originalText := message.ReplyTo.Text

	msg := domain.NewTrainMessageWithReply(chatId, originalText, replyText)
	a.learn.Push(msg)
}
