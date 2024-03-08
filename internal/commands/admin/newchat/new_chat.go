package newchat

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/repository"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	ChatID       int64
	Status       string
	FromUserName string
}

func Do(bot *tgbotapi.BotAPI, s3Repo *repository.S3Repository, dto DTO) *logger.SlogError {
	if dto.Status != "member" {
		if dto.Status != "left" {
			slog.Warn("got unknown status for new chat", slog.String("status", dto.Status))
		}
		return nil
	}
	if dto.FromUserName != AdminUserName {
		_, err := bot.Request(tgbotapi.LeaveChatConfig{
			ChatID: dto.ChatID,
		})
		if err != nil {
			return logger.NewSlogError(err, "error leaving chat")
		}
		return nil
	}

	cfg, err := s3Repo.GetConfig()
	if err != nil {
		return err
	}
	cfg.AdminChatID = dto.ChatID
	if err := s3Repo.SaveConfig(*cfg); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(
		dto.ChatID,
		"Чат для получения уведомлений обновлен",
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
