package accept

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/shared"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	ChatID int64
	MsgID  int
}

func Do(bot *tgbotapi.BotAPI, dto DTO) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		dto.ChatID,
		"Вы приняли приглашение. Мы рады, что вы будете с нами!",
	)
	msg.ReplyMarkup = shared.GetFinishReplyKeyboard()
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
