package wishes

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/notifications"
	"github.com/gonozov0/weddingtgbot/internal/repository/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	TgID   int64
	ChatID int64
	Wishes string
}

func Do(bot *tgbotapi.BotAPI, s3Repo *s3.Repository, dto DTO) *logger.SlogError {
	anws, err := s3Repo.GetAnswers(dto.TgID)
	if err != nil {
		return err
	}
	anws.Wishes = dto.Wishes
	if err = s3Repo.SaveAnswers(*anws); err != nil {
		return err
	}

	cfg, err := s3Repo.GetConfig()
	if err != nil {
		return err
	}
	if err = notifications.SendComplete(bot, cfg.AdminChatID, anws.FirstName); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(dto.ChatID, completeMessage)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
