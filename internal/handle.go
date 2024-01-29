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
		if update.Message.Contact != nil {
			return commands.Start(bot, commands.StartDTO{
				ChatID: update.Message.Chat.ID,
				Login:  update.Message.Contact.PhoneNumber,
			})
		}

		switch update.Message.Text {
		case commands.StartCommand, commands.StartAgainCommand:
			return commands.Start(bot, commands.StartDTO{
				ChatID: update.Message.Chat.ID,
				Login:  update.Message.From.UserName,
			})
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
		default:
			return commands.Unknown(bot, update.Message.Chat.ID)
		}
	}

	return logger.NewSlogError(nil, "got unknown update type")
}
