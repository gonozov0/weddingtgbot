package withsomebody

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/shared"
	"github.com/gonozov0/weddingtgbot/internal/repository/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	TgID   int64
	ChatID int64
}

func Do(bot *tgbotapi.BotAPI, s3Repo *s3.Repository, dto DTO) *logger.SlogError {
	anws, err := s3Repo.GetAnswers(dto.TgID)
	if err != nil {
		return err
	}
	withSomebody := true
	anws.WithSomebody = &withSomebody
	err = s3Repo.SaveAnswers(*anws)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(dto.ChatID, shared.SecondGuestMessage)
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, Selective: true}
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
