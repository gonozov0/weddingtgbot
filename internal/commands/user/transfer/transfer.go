package transfer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/shared"
	"github.com/gonozov0/weddingtgbot/internal/repository/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	TgID    int64
	ChatID  int64
	Command string
}

func Do(bot *tgbotapi.BotAPI, s3Repo *s3.Repository, dto DTO) *logger.SlogError {
	anws, err := s3Repo.GetAnswers(dto.TgID)
	if err != nil {
		return err
	}

	var needTransfer string
	switch dto.Command {
	case commands.TransferNotNeeded:
		needTransfer = "No"
	case commands.RostovTransferNeeded:
		needTransfer = "Rostov"
	case commands.YaroslavlTransferNeeded:
		needTransfer = "Yaroslavl"
	}
	anws.NeedTransfer = needTransfer

	err = s3Repo.SaveAnswers(*anws)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(dto.ChatID, shared.WishesMessage)
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, Selective: true}
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
