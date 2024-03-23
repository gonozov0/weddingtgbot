package sendanalytic

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/repository/googledoc"
	"github.com/gonozov0/weddingtgbot/internal/repository/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func Do(bot *tgbotapi.BotAPI, s3Repo *s3.Repository, gglRepo *googledoc.Repository, chatID int64) *logger.SlogError {
	guestAnswers, err := s3Repo.GetAllAnswers()
	if err != nil {
		return err
	}

	answers := make([]googledoc.AnswerDTO, 0, len(guestAnswers))
	for _, ga := range guestAnswers {
		answers = append(answers, googledoc.AnswerDTO{
			FirstName:    ga.FirstName,
			LastName:     ga.LastName,
			IsAccepted:   ga.IsAccepted,
			WithSomebody: ga.WithSomebody,
			SecondGuest:  ga.SecondGuest,
			NeedTransfer: ga.NeedTransfer,
			Wishes:       ga.Wishes,
		})
	}

	err = gglRepo.InsertAnswers(answers)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, "Аналитика отправлена")
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
