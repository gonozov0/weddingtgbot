package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

type AcceptDTO struct {
	ChatID int64
	MsgID  int
}

func Accept(bot *tgbotapi.BotAPI, dto AcceptDTO) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		dto.ChatID,
		"Вы приняли приглашение. Мы рады, что вы будете с нами!",
	)
	msg.ReplyMarkup = getFinishReplyKeyboard()
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
