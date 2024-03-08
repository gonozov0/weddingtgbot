package user

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/accept"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/alone"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/decline"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/secondguest"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/shared"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/start"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/transfer"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/unknown"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/wishes"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/withsomebody"
	"github.com/gonozov0/weddingtgbot/internal/repository"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func HandleCommands(bot *tgbotapi.BotAPI, s3Repo *repository.S3Repository, update tgbotapi.Update) *logger.SlogError {
	if update.Message.ReplyToMessage != nil {
		if strings.TrimSpace(update.Message.ReplyToMessage.Text) == strings.TrimSpace(shared.SecondGuestMessage) {
			return secondguest.Do(bot, s3Repo, secondguest.DTO{
				TgID:        update.Message.From.ID,
				ChatID:      update.Message.Chat.ID,
				SecondGuest: update.Message.Text,
			})
		}
		if strings.TrimSpace(update.Message.ReplyToMessage.Text) == strings.TrimSpace(shared.WishesMessage) {
			return wishes.Do(bot, s3Repo, wishes.DTO{
				TgID:   update.Message.From.ID,
				ChatID: update.Message.Chat.ID,
				Wishes: update.Message.Text,
			})
		}
	}

	switch update.Message.Text {
	case commands.Start:
		return start.Do(bot, s3Repo, start.DTO{
			TgID:   update.Message.From.ID,
			ChatID: update.Message.Chat.ID,
		})
	case commands.Accept:
		return accept.Do(bot, s3Repo, accept.DTO{
			TgID:   update.Message.From.ID,
			ChatID: update.Message.Chat.ID,
			MsgID:  update.Message.MessageID,
		})
	case commands.Decline:
		return decline.Do(bot, s3Repo, decline.DTO{
			TgID:   update.Message.From.ID,
			ChatID: update.Message.Chat.ID,
			MsgID:  update.Message.MessageID,
		})
	case commands.Alone:
		return alone.Do(bot, s3Repo, alone.DTO{
			TgID:   update.Message.From.ID,
			ChatID: update.Message.Chat.ID,
		})
	case commands.WithSomebody:
		return withsomebody.Do(bot, s3Repo, withsomebody.DTO{
			TgID:   update.Message.From.ID,
			ChatID: update.Message.Chat.ID,
		})
	case commands.TransferNotNeeded, commands.RostovTransferNeeded, commands.YaroslavlTransferNeeded:
		return transfer.Do(bot, s3Repo, transfer.DTO{
			TgID:    update.Message.From.ID,
			ChatID:  update.Message.Chat.ID,
			Command: update.Message.Text,
		})
	}

	return unknown.Do(bot, update.Message.Chat.ID)
}
