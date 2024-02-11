package handle_new_chat

import (
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	ChatID int64
	Status string
}

func Do(bot *tgbotapi.BotAPI, dto DTO) *logger.SlogError {
	if dto.Status != "member" {
		slog.Warn("got unknown status for new chat", slog.String("status", dto.Status))
		return nil
	}

	msg := tgbotapi.NewMessage(
		dto.ChatID,
		fmt.Sprintf("ID этого чата: %d", dto.ChatID),
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending photo file id")
	}

	return nil
}
