package unknown

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func Do(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		chatID,
		"Выберите, пожалуйста, один из предложенных вариантов в интерфейсе",
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
