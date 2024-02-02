package main

import (
	"log"
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

func main() {
	logger.Setup()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Webhook successfully deleted", slog.String("bot", bot.Self.UserName))
}
