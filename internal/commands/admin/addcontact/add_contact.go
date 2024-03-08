package addcontact

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/repository"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	TgID   int64
	ChatID int64
	Name   string
}

func Do(bot *tgbotapi.BotAPI, s3Repo *repository.S3Repository, dto DTO) *logger.SlogError {
	config, err := s3Repo.GetConfig()
	if err != nil {
		return err
	}

	config.GuestsInfo[repository.TgID(dto.TgID)] = repository.Name(dto.Name)
	if err := s3Repo.SaveConfig(*config); err != nil {
		return err
	}

	if _, err := bot.Send(tgbotapi.NewMessage(dto.ChatID, "Contact added to guests' info")); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
