package owner_chat

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

const declineMessage = "Гость не сможет присутствовать на свадьбе :("

func SendDecline(bot *tgbotapi.BotAPI) *logger.SlogError {
	msg := tgbotapi.NewMessage(chatID, declineMessage)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}
	return nil
}
