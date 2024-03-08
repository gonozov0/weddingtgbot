package decline

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/shared"
	"github.com/gonozov0/weddingtgbot/internal/repository"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	TgID   int64
	ChatID int64
	MsgID  int
}

func Do(bot *tgbotapi.BotAPI, s3Repo *repository.S3Repository, dto DTO) *logger.SlogError {
	anws, err := s3Repo.GetAnswers(dto.TgID)
	if err != nil {
		return err
	}
	accepted := false
	anws.IsAccepted = &accepted
	err = s3Repo.SaveAnswers(*anws)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(
		dto.ChatID,
		declineText,
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	contact := shared.GetOlyaContact(dto.ChatID)
	contact.ReplyMarkup = shared.GetEmptyReplyKeyboard()
	if _, err := bot.Send(contact); err != nil {
		return logger.NewSlogError(err, "error sending contact")
	}

	return nil
}
