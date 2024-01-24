package internal

import (
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/internal/internal/callbacks"
	"github.com/gonozov0/weddingbot/internal/internal/commands"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) *logger.SlogError {
	if update.Message != nil {
		if update.Message.Text == commands.StartCommand {
			return commands.Start(bot, update.Message.Chat.ID)
		}

		return commands.Unknown(bot, update.Message.Chat.ID)
	}

	if update.CallbackQuery != nil {
		switch update.CallbackQuery.Data {
		case callbacks.AcceptCallback:
			return callbacks.Accept(bot, callbacks.AcceptDTO{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				MsgID:  update.CallbackQuery.Message.MessageID,
			})
		case callbacks.DeclineCallback:
			return callbacks.Decline(bot, callbacks.DeclineDTO{
				ChatID: update.CallbackQuery.Message.Chat.ID,
				MsgID:  update.CallbackQuery.Message.MessageID,
			})
		default:
			return logger.NewSlogError(
				nil,
				"unknown callback query data",
				slog.String("data", update.CallbackQuery.Data),
			)
		}
	}

	return logger.NewSlogError(nil, "got unknown update type")
}
