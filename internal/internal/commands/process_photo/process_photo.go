package process_photo

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	ChatID int64
	Photo  []tgbotapi.PhotoSize
}

func Do(bot *tgbotapi.BotAPI, dto DTO) *logger.SlogError {
	if _, err := bot.Send(tgbotapi.NewMessage(dto.ChatID, dto.Photo[0].FileID)); err != nil {
		return logger.NewSlogError(err, "error sending photo file id")
	}

	return nil
}
