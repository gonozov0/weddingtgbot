package owner_chat

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

const withSomebodyMessage = "Гость придет с кем-то"

func SendWithSomebody(bot *tgbotapi.BotAPI) *logger.SlogError {
	msg := tgbotapi.NewMessage(chatID, withSomebodyMessage)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}
	return nil
}
