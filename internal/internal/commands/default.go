package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

func Unknown(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		chatID,
		"Извините, я вас не понимаю. Пожалуйста, выберите один из вариантов ответа выше.",
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
