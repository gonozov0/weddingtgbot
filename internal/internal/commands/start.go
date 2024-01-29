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

type StartDTO struct {
	ChatID int64
	Login  string
}

func Start(bot *tgbotapi.BotAPI, dto StartDTO) *logger.SlogError {
	if dto.Login == "" {
		return requestPhoneNumber(bot, dto.ChatID)
	}

	if !isLoginInvited(dto.Login) {
		return sendNotInvitedInfo(bot, dto.ChatID)
	}

	return sendInvitation(bot, dto.ChatID)
}

func requestPhoneNumber(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		chatID,
		"Пожалуйста, предоставьте свой номер телефона для проверки, что вы есть в списке гостей.",
	)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("Отправить мой номер боту для идентификации"),
		),
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}

var loginWhitelist = []string{
	"gonozov0",
	"TaoGen",
	"+7 915 979 6484",
}

func isLoginInvited(login string) bool {
	for _, whitelistedLogin := range loginWhitelist {
		if login == whitelistedLogin {
			return true
		}
	}
	return false
}

func sendNotInvitedInfo(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	msg := tgbotapi.NewMessage(
		chatID,
		"К сожалению, вас нет в списке приглашенных. Если вы считаете, что это ошибка, пожалуйста, свяжитесь с нами.",
	)
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	contact := tgbotapi.NewContact(chatID, "+79807442720", "Оля")
	contact.ReplyMarkup = getFinishReplyKeyboard()
	if _, err := bot.Send(contact); err != nil {
		return logger.NewSlogError(err, "error sending contact")
	}

	return nil
}

func sendInvitation(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	photoGroup := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(photoFileID))
	if _, err := bot.Send(photoGroup); err != nil {
		return logger.NewSlogError(err, "error sending photo")
	}

	newMsg := tgbotapi.NewMessage(chatID, invitationText)
	newMsg.ParseMode = "Markdown"
	newMsg.ReplyMarkup = getStartReplyKeyboard()
	if _, err := bot.Send(newMsg); err != nil {
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
