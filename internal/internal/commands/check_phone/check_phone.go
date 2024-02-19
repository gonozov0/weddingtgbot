package check_phone

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/shared"
	"github.com/gonozov0/weddingtgbot/internal/internal/commands/shared/owner_chat"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
	"github.com/gonozov0/weddingtgbot/pkg/phone_utils"
)

type DTO struct {
	ChatID int64
	Phone  string
}

func Do(bot *tgbotapi.BotAPI, dto DTO) *logger.SlogError {
	if !isPhoneInvited(dto.Phone) {
		return shared.SendNotInvitedInfo(bot, dto.ChatID)
	}

	personInfo := shared.GetPersonInfo(phone_utils.Normalize(dto.Phone))
	if err := owner_chat.SendStart(bot, personInfo.GetFullName()); err != nil {
		return err
	}

	return sendInvitation(bot, dto.ChatID, dto.Phone)
}

func sendInvitation(bot *tgbotapi.BotAPI, chatID int64, name string) *logger.SlogError {
	newMsg := tgbotapi.NewMessage(chatID, fmt.Sprintf(invitationGuestText, name))
	newMsg.ParseMode = "Markdown"
	newMsg.ReplyMarkup = shared.GetStartReplyKeyboard()
	if _, err := bot.Send(newMsg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}
