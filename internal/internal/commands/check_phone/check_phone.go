package check_phone

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/shared"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	ChatID int64
	Phone  string
}

func Do(bot *tgbotapi.BotAPI, dto DTO) *logger.SlogError {
	if !isPhoneInvited(dto.Phone) {
		return shared.SendNotInvitedInfo(bot, dto.ChatID)
	}

	return sendInvitation(bot, dto.ChatID)
}

func sendInvitation(bot *tgbotapi.BotAPI, chatID int64) *logger.SlogError {
	newMsg := tgbotapi.NewMessage(chatID, invitationGuestText)
	newMsg.ParseMode = "Markdown"
	newMsg.ReplyMarkup = shared.GetAnswerReplyKeyboard()
	if _, err := bot.Send(newMsg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
