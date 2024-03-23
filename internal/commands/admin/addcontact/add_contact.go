package addcontact

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/repository/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	TgID      int64
	ChatID    int64
	FirstName string
	LastName  string
}

func Do(bot *tgbotapi.BotAPI, s3Repo *s3.Repository, dto DTO) *logger.SlogError {
	config, err := s3Repo.GetConfig()
	if err != nil {
		return err
	}

	config.GuestsInfo[s3.TgID(dto.TgID)] = s3.GuestInfo{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
	if err := s3Repo.SaveConfig(*config); err != nil {
		return err
	}

	if _, err := bot.Send(tgbotapi.NewMessage(dto.ChatID, "Contact added to guests' info")); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
