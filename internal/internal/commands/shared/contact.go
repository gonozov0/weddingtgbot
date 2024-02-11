package shared

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetOlyaContact(chatID int64) tgbotapi.ContactConfig {
	return tgbotapi.NewContact(chatID, "+79807442720", "Оля")
}
