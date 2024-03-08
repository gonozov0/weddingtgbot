package notifications

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

const completeMessage = "Гость %s завершил ответы на вопросы"

func SendComplete(bot *tgbotapi.BotAPI, ownerChatID int64, name string) *logger.SlogError {
	msg := tgbotapi.NewMessage(ownerChatID, fmt.Sprintf(completeMessage, name))
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
