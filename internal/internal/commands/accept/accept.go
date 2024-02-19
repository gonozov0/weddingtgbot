package accept

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/shared/owner_chat"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	ChatID int64
	MsgID  int
}

func Do(bot *tgbotapi.BotAPI, dto DTO) *logger.SlogError {
	if err := owner_chat.SendAccept(bot); err != nil {
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
