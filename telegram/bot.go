package telegram

import (
	"log"
	"net/http"
	"time"

	"github.com/Skoowshot/vecspect/domain"
)

type TelegramBot struct {
	token       string
	httpClient  http.Client
	pollTimeout int
}

type UpdateListener interface {
	OnUpdate(update *domain.Update)
}

func NewTelegramBot(token string, pollTimeout int) *TelegramBot {
	clientTimeout := time.Duration(pollTimeout+5) * time.Second

	return &TelegramBot{
		token: token,
		httpClient: http.Client{
			Timeout: clientTimeout,
		},
		pollTimeout: pollTimeout,
	}
}

func (b *TelegramBot) PollUpdates(callback UpdateListener) {
	offset := 0

	for {
		updates, err := b.GetUpdates(offset)
		if err != nil {
			log.Println("error while getting updates:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, update := range updates.Result {
			callback.OnUpdate(&update)

			offset = int(update.UpdateId + 1)
		}
	}
}
