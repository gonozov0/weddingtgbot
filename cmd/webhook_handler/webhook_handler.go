package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gonozov0/weddingtgbot/internal"
	"github.com/gonozov0/weddingtgbot/internal/repository/googledoc"
	"github.com/gonozov0/weddingtgbot/internal/repository/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
	"github.com/gonozov0/weddingtgbot/pkg/yandex_cloud/lambda"
)

func Handler(ctx context.Context, rawReq []byte) (*lambda.Response, error) {
	logger.Setup()

	body, err := lambda.DecodeBody(rawReq)
	if err != nil {
		return nil, err
	}

	var update tgbotapi.Update
	err = json.Unmarshal(body, &update)
	if err != nil {
		slog.Error("Error unmarshalling update", slog.String("err", err.Error()))
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		slog.Error("Error initializing bot", slog.String("err", err.Error()))
		return nil, err
	}

	s3Repo, slogErr := s3.NewRepository()
	if slogErr != nil {
		slogErr.Log(ctx)
		return nil, slogErr
	}

	gglRepo, slogErr := googledoc.NewRepository(
		"creds/avid-sphere-413708-c20f667d42e2.json",
		"1csSzXWCeDkDZirCwae4ozG0i0PnMmT42dttnmq8iDMo",
	)
	if slogErr != nil {
		slogErr.Log(ctx)
		return nil, slogErr
	}

	slogErr = internal.HandleUpdate(bot, s3Repo, gglRepo, update)
	if slogErr != nil {
		slogErr.Log(ctx)
		return nil, slogErr
	}

	return &lambda.Response{
		StatusCode: http.StatusOK,
		Body:       "ok",
	}, nil
}
