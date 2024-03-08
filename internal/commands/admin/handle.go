package admin

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands/admin/addcontact"
	"github.com/gonozov0/weddingtgbot/internal/commands/admin/addphoto"
	"github.com/gonozov0/weddingtgbot/internal/commands/admin/newchat"
	"github.com/gonozov0/weddingtgbot/internal/repository"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func HandleCommands(bot *tgbotapi.BotAPI, s3Repo *repository.S3Repository, update tgbotapi.Update) *logger.SlogError {
	if update.MyChatMember != nil {
		return newchat.Do(bot, s3Repo, newchat.DTO{
			ChatID:       update.MyChatMember.Chat.ID,
			Status:       update.MyChatMember.NewChatMember.Status,
			FromUserName: update.MyChatMember.From.UserName,
		})
	}
	if update.Message.Photo != nil {
		return addphoto.Do(bot, s3Repo, addphoto.DTO{
			ChatID:      update.Message.Chat.ID,
			PhotoFileID: update.Message.Photo[0].FileID,
		})
	}
	if update.Message.Contact != nil {
		return addcontact.Do(bot, s3Repo, addcontact.DTO{
			TgID:   update.Message.Contact.UserID,
			ChatID: update.Message.Chat.ID,
			Name:   update.Message.Contact.FirstName,
		})
	}

	return logger.NewSlogError(nil, "got unknown admin command")
}
