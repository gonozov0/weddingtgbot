package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

func Start(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		chatID,
		"Здравствуйте, вы приглашены на свадьбу Гонозовых, которая состоится ...",
	)
	msg.ReplyMarkup = getStartingInlineKeyboard()
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}

func getStartingInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Принять", "accept"),
			tgbotapi.NewInlineKeyboardButtonData("Отказаться", "decline"),
		),
	)
}
