package start

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal/commands/user/shared"
	"github.com/gonozov0/weddingtgbot/internal/notifications"
	"github.com/gonozov0/weddingtgbot/internal/repository/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type DTO struct {
	TgID   int64
	ChatID int64
}

func Do(bot *tgbotapi.BotAPI, s3Repo *s3.Repository, dto DTO) *logger.SlogError {
	config, err := s3Repo.GetConfig()
	if err != nil {
		return err
	}

	gi, isInvited := config.GuestsInfo[s3.TgID(dto.TgID)]
	if !isInvited {
		return sendNotInvitedInfo(bot, dto.ChatID, config.PhotoFileID)
	}

	err = s3Repo.SaveAnswers(s3.GuestAnswers{
		TgID:      dto.TgID,
		FirstName: gi.FirstName,
		LastName:  gi.LastName,
	})
	if err != nil {
		return err
	}

	err = notifications.SendStart(bot, config.AdminChatID, gi.FirstName, gi.LastName)
	if err != nil {
		return err
	}

	return sendInvitation(bot, dto.ChatID, gi.FirstName, config.PhotoFileID)
}

func sendInvitation(bot *tgbotapi.BotAPI, chatID int64, name string, photoFileID string) *logger.SlogError {
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(photoFileID))
	if _, err := bot.Send(photo); err != nil {
		return logger.NewSlogError(err, "error sending photo")
	}

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf(invitationGuestText, name))
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = shared.GetStartReplyKeyboard()
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	return nil
}

func sendNotInvitedInfo(bot *tgbotapi.BotAPI, chatID int64, photoFileID string) *logger.SlogError {
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileID(photoFileID))
	if _, err := bot.Send(photo); err != nil {
		return logger.NewSlogError(err, "error sending photo")
	}

	msg := tgbotapi.NewMessage(
		chatID,
		notInvitedText,
	)
	msg.ParseMode = "Markdown"
	if _, err := bot.Send(msg); err != nil {
		return logger.NewSlogError(err, "error sending message")
	}

	contact := shared.GetOlyaContact(chatID)
	contact.ReplyMarkup = shared.GetEmptyReplyKeyboard()
	if _, err := bot.Send(contact); err != nil {
		return logger.NewSlogError(err, "error sending contact")
	}

	return nil
}
