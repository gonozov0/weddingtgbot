package admin

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands/admin/addcontact"
	"github.com/gonozov0/weddingtgbot/internal/commands/admin/addphoto"
	"github.com/gonozov0/weddingtgbot/internal/commands/admin/newchat"
	"github.com/gonozov0/weddingtgbot/internal/commands/admin/sendanalytic"
	"github.com/gonozov0/weddingtgbot/internal/repository/googledoc"
	"github.com/gonozov0/weddingtgbot/internal/repository/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func HandleCommands(
	bot *tgbotapi.BotAPI,
	s3Repo *s3.Repository,
	gglRepo *googledoc.Repository,
	update tgbotapi.Update,
) *logger.SlogError {
	if update.MyChatMember != nil {
		return newchat.Do(bot, s3Repo, newchat.DTO{
			ChatID:    update.MyChatMember.Chat.ID,
			Status:    update.MyChatMember.NewChatMember.Status,
			FromAdmin: update.MyChatMember.From.UserName == adminUserName,
		})
	}
	if update.Message.From.UserName == adminUserName {
		if update.Message.Text == SendAnalyticCommand {
			return sendanalytic.Do(bot, s3Repo, gglRepo, update.Message.Chat.ID)
		}
		if update.Message.Photo != nil {
			return addphoto.Do(bot, s3Repo, addphoto.DTO{
				ChatID:      update.Message.Chat.ID,
				PhotoFileID: update.Message.Photo[0].FileID,
			})
		}
		if update.Message.Contact != nil {
			return addcontact.Do(bot, s3Repo, addcontact.DTO{
				TgID:      update.Message.Contact.UserID,
				ChatID:    update.Message.Chat.ID,
				FirstName: update.Message.Contact.FirstName,
				LastName:  update.Message.Contact.LastName,
			})
		}
	}

	return sendUnknown(bot, update.Message.Chat.ID)
}

func sendUnknown(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	msg := tgbotapi.NewMessage(chatID, "Неизвестная команда")
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}
	return nil
}
