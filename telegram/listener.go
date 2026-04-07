package telegram

import "github.com/Skoowshot/vecspect/domain"

type DefaultUpdateListener struct {
	OnMessage func(message *domain.Message)
}

func (l *DefaultUpdateListener) OnUpdate(update *domain.Update) {
	if msg := update.Message; msg != nil {
		l.OnMessage(update.Message)
	}
}

func NewDefaultUpdateListener(onMessage func(message *domain.Message)) *DefaultUpdateListener {
	return &DefaultUpdateListener{
		OnMessage: onMessage,
	}
}
