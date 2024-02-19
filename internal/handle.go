package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/accept"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/alone"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/check_phone"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/decline"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/handle_new_chat"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/process_photo"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/start"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/unknown"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/with_somebody"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) *logger.SlogError {
	if update.MyChatMember != nil {
		return handle_new_chat.Do(bot, handle_new_chat.DTO{
			ChatID: update.MyChatMember.Chat.ID,
			Status: update.MyChatMember.NewChatMember.Status,
		})
	}

	if update.Message != nil {
		if update.Message.From.ID == bot.Self.ID ||
			update.Message.NewChatMembers != nil ||
			update.Message.LeftChatMember != nil ||
			update.Message.GroupChatCreated ||
			update.Message.SuperGroupChatCreated ||
			update.Message.ChannelChatCreated ||
			update.Message.MigrateToChatID != 0 ||
			update.Message.MigrateFromChatID != 0 {
			return nil
		}

		if update.Message.Photo != nil {
			return process_photo.Do(bot, process_photo.DTO{
				ChatID: update.Message.Chat.ID,
				Photo:  update.Message.Photo,
			})
		}
		if update.Message.Contact != nil {
			return check_phone.Do(bot, check_phone.DTO{
				ChatID: update.Message.Chat.ID,
				Phone:  update.Message.Contact.PhoneNumber,
			})
		}

		switch update.Message.Text {
		case commands.Start:
			return start.Do(bot, start.DTO{
				ChatID: update.Message.Chat.ID,
				Login:  update.Message.From.UserName,
			})
		case commands.Accept:
			return accept.Do(bot, accept.DTO{
				ChatID: update.Message.Chat.ID,
				MsgID:  update.Message.MessageID,
			})
		case commands.Decline:
			return decline.Do(bot, decline.DTO{
				ChatID: update.Message.Chat.ID,
				MsgID:  update.Message.MessageID,
			})
		case commands.Alone:
			return alone.Do(bot, update.Message.Chat.ID)
		case commands.WithSomebody:
			return with_somebody.Do(bot, update.Message.Chat.ID)
		default:
			return unknown.Do(bot, update.Message.Chat.ID)
		}
	}

	return logger.NewSlogError(nil, "got unknown update type")
}
