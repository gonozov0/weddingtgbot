package shared

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands"
)

func GetStartReplyKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(commands.Accept),
			tgbotapi.NewKeyboardButton(commands.Decline),
		),
	)
	keyboard.OneTimeKeyboard = false
	return keyboard
}
