package accept

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands"
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
	accepted := true
	anws.IsAccepted = &accepted
	err = s3Repo.SaveAnswers(*anws)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(
		dto.ChatID,
		acceptText,
	)
	msg.ReplyMarkup = getReplyKeyboard()
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}

func getReplyKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(commands.Alone),
			tgbotapi.NewKeyboardButton(commands.WithSomebody),
		),
	)
}
