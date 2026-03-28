package telegram

import (
	"net/http"
	"time"
)

type TelegramBot struct {
	token       string
	httpClient  http.Client
	pollTimeout int
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
