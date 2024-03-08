package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands/admin"
	"github.com/gonozov0/weddingtgbot/internal/commands/user"
	"github.com/gonozov0/weddingtgbot/internal/repository"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func HandleUpdate(bot *tgbotapi.BotAPI, s3Repo *repository.S3Repository, update tgbotapi.Update) *logger.SlogError {
	if update.MyChatMember == nil && update.Message == nil {
		return logger.NewSlogError(nil, "got unknown update command")
	}

	if excludeHandling(update, bot.Self.ID) {
		return nil
	}

	if update.MyChatMember != nil || update.Message.Photo != nil || update.Message.Contact != nil {
		return admin.HandleCommands(bot, s3Repo, update)
	}

	return user.HandleCommands(bot, s3Repo, update)
}

func excludeHandling(update tgbotapi.Update, botID int64) bool {
	return update.Message != nil && (update.Message.From.ID == botID ||
		update.Message.NewChatMembers != nil ||
		update.Message.LeftChatMember != nil ||
		update.Message.GroupChatCreated ||
		update.Message.SuperGroupChatCreated ||
		update.Message.ChannelChatCreated ||
		update.Message.MigrateToChatID != 0 ||
		update.Message.MigrateFromChatID != 0)
}
