package shared

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands"
)

const TransferMessage = `
Так как место проведения далеко от ближайших городов, мы организуем трансфер после завершения мероприятия.
Скорее всего это будет маршрутка.
Ответьте, пожалуйста, нужен ли вам трансфер?
`

func GetTransferReplyKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(commands.TransferNotNeeded),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(commands.RostovTransferNeeded),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(commands.YaroslavlTransferNeeded),
		),
	)
}
