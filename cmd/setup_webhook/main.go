package main

import (
	"log"
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingbot/pkg/logger"
)

func main() {
	logger.Setup()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	webhookURL := os.Getenv("WEBHOOK_URL")
	webhookConfig, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = bot.Request(webhookConfig)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Web hook set up", slog.String("url", webhookURL), slog.String("bot", bot.Self.UserName))
}
