package s3

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

const botConfigPath = "bot_config.json"

type (
	TgID int64
)

type GuestInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type BotConfig struct {
	PhotoFileID string             `json:"photo_file_id"`
	AdminChatID int64              `json:"admin_chat_id"`
	GuestsInfo  map[TgID]GuestInfo `json:"guests_info"`
}

func (repo *Repository) GetConfig() (*BotConfig, *logger.SlogError) {
	resp, err := repo.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &repo.bucketName,
		Key:    aws.String(botConfigPath),
	})
	if err != nil {
		return nil, logger.NewSlogError(err, "error getting object from s3")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, logger.NewSlogError(err, "error reading object body")
	}

	var config *BotConfig
	if err := json.Unmarshal(body, &config); err != nil {
		return nil, logger.NewSlogError(err, "error unmarshalling bot config")
	}

	return config, nil
}

func (repo *Repository) SaveConfig(config BotConfig) *logger.SlogError {
	body, err := json.Marshal(config)
	if err != nil {
		return logger.NewSlogError(err, "error marshalling bot config")
	}

	_, err = repo.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &repo.bucketName,
		Key:    aws.String(botConfigPath),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		return logger.NewSlogError(err, "error putting object to s3")
	}

	return nil
}
