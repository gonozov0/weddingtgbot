package with_somebody

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/shared"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/shared/owner_chat"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func Do(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	if err := owner_chat.SendWithSomebody(bot); err != nil {
		return err
	}

	// TODO: поменять на логику получения списка дополнительных гостей
	msg := tgbotapi.NewMessage(chatID, shared.TransferMessage)
	msg.ReplyMarkup = shared.GetTransferReplyKeyboard()
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}
	return nil
}
