package notifications

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

const startMessage = "Гость %s %s начал диалог с ботом"

func SendStart(bot *tgbotapi.BotAPI, adminChatID int64, firstName, lastName string) *logger.SlogError {
	msg := tgbotapi.NewMessage(adminChatID, fmt.Sprintf(startMessage, firstName, lastName))
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
