package decline

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/shared"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

const (
	declineAnswer = `
Очень жаль, что вы не сможете быть с нами :(
Если вдруг ваше решение изменится, напишите Оле, чтобы сообщить эту радостную новость:
`
)

type DTO struct {
	ChatID int64
	MsgID  int
}

func Do(bot *tgbotapi.BotAPI, dto DTO) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		dto.ChatID,
		declineAnswer,
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	contact := shared.GetOlyaContact(dto.ChatID)
	contact.ReplyMarkup = shared.GetFinishReplyKeyboard()
	if _, err := bot.Send(contact); err != nil {
		return logger.NewSlogError(err, "error sending contact")
	}

	return nil
}
