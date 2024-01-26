package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/internal/internal/commands"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) *logger.SlogError {
	if update.Message != nil {
		if update.Message.Photo != nil {
			return commands.Photo(bot, commands.PhotoDTO{
				ChatID: update.Message.Chat.ID,
				Photo:  update.Message.Photo,
			})
		}

		switch update.Message.Text {
		case commands.StartCommand, commands.ChangeAnswerCommand:
			return commands.Start(bot, update.Message.Chat.ID)
		case commands.AcceptCommand:
			return commands.Accept(bot, commands.AcceptDTO{
				ChatID: update.Message.Chat.ID,
				MsgID:  update.Message.MessageID,
			})
		case commands.DeclineCommand:
			return commands.Decline(bot, commands.DeclineDTO{
				ChatID: update.Message.Chat.ID,
				MsgID:  update.Message.MessageID,
			})
		}

		return commands.Unknown(bot, update.Message.Chat.ID)
	}

	return logger.NewSlogError(nil, "got unknown update type")
}
