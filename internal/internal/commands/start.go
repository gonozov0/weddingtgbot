package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

const (
	photoFileID    = "AgACAgIAAxkBAAOBZbOw1UdDCWJZLJqd-djGgHgqaIoAAvn1MRvzbZhJrgHQueo7tmkBAAMCAAN5AAM0BA"
	invitationText = `
*Приглашение на Свадьбу!*

Дорогие друзья,

Мы, _Дима и Оля Гонозовы_, рады пригласить вас на торжество по случаю нашего бракосочетания!

📅 *Дата:* 24 июля 2024 года
📍 *Место проведения:* Ресторан "Романтик", г. Москва

Ваше присутствие будет для нас лучшим подарком!

С любовью,
*Дима и Оля*
`
)

func Start(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	photoGroup := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(photoFileID))
	if _, err := bot.Send(photoGroup); err != nil {
		return logger.NewSlogError(err, "error sending photo")
	}

	msg := tgbotapi.NewMessage(chatID, invitationText)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = getStartReplyKeyboard()
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}

func getStartReplyKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(AcceptCommand),
			tgbotapi.NewKeyboardButton(DeclineCommand),
		),
	)
}
