package domain

type TrainingMessage struct {
	ChatId   int64
	Original string
	Reply    *string
}

func NewTrainMessageWithReply(chatId int64, original string, reply string) TrainingMessage {
	return TrainingMessage{
		ChatId:   chatId,
		Original: original,
		Reply:    &reply,
	}
}

func NewTrainMessage(chatId int64, original string) TrainingMessage {
	return TrainingMessage{
		ChatId:   chatId,
		Original: original,
	}
}
