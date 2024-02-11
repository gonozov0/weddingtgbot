package shared

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

const notInvitedText = `
К сожалению, вас нет в списке приглашенных.
Если вы не сами нашли этого бота, то напишите, пожалуйста, Оле, чтобы уточнить информацию:
`

func SendNotInvitedInfo(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		chatID,
		notInvitedText,
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	contact := GetOlyaContact(chatID)
	contact.ReplyMarkup = GetFinishReplyKeyboard()
	if _, err := bot.Send(contact); err != nil {
		return logger.NewSlogError(err, "error sending contact")
	}

	return nil
}
