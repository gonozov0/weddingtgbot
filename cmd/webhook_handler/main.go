package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/pkg/logger"
	"github.com/gonozov0/weddingbot/pkg/yandex_cloud/lambda"
)

func Handler(ctx context.Context, rawReq []byte) (*lambda.Response, error) {
	logger.Setup()

	body, err := lambda.DecodeBody(rawReq)
	if err != nil {
		return nil, err
	}

	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		slog.Error("Error initializing bot", slog.String("err", err.Error()))
		return nil, err
	}

	if update.Message != nil && update.Message.Text != "" {
		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Здравствуйте, вы приглашены на свадьбу Гонозовых, которая состоится ...",
			)
			msg.ReplyMarkup = getInlineKeyboard()
			if _, err := bot.Send(msg); err != nil {
				slog.Error("Error sending message", slog.String("err", err.Error()))
				return nil, err
			}
		}
	}

	if update.CallbackQuery != nil {
		var callbackText string
		switch update.CallbackQuery.Data {
		case "accept":
			callbackText = "Вы приняли приглашение. Мы рады, что вы будете с нами!"
		case "decline":
			callbackText = "Вы отказались от приглашения. Мы будем скучать!"
		}
		callbackConfig := tgbotapi.NewCallback(update.CallbackQuery.ID, callbackText)
		if _, err := bot.Request(callbackConfig); err != nil {
			slog.Error("Error sending callback", slog.String("err", err.Error()))
			return nil, err
		}
	}

	return &lambda.Response{
		StatusCode: 200,
		Body:       "ok",
	}, nil
}

func getInlineKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Принять", "accept"),
			tgbotapi.NewInlineKeyboardButtonData("Отказаться", "decline"),
		),
	)
}
