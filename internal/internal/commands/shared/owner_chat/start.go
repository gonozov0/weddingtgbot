package owner_chat

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

const startMessage = "Гость %s начал диалог с ботом"

func SendStart(bot *tgbotapi.BotAPI, fullName string) *logger.SlogError {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(startMessage, fullName))
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}
	return nil
}
