package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

type DeclineDTO struct {
	ChatID int64
	MsgID  int
}

func Decline(bot *tgbotapi.BotAPI, dto DeclineDTO) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		dto.ChatID,
		"Вы отказались от приглашения. Мы будем скучать!",
	)
	msg.ReplyMarkup = getFinishReplyKeyboard()
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
