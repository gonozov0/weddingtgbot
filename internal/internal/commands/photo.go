package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

type PhotoDTO struct {
	ChatID int64
	Photo  []tgbotapi.PhotoSize
}

func Photo(bot *tgbotapi.BotAPI, dto PhotoDTO) *logger.SlogError {
	for _, photo := range dto.Photo {
		if _, err := bot.Send(tgbotapi.NewMessage(dto.ChatID, photo.FileID)); err != nil {
			return logger.NewSlogError(err, "error sending message")
		}
	}

	return nil
}
