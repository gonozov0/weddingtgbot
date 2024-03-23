package s3

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type GuestAnswers struct {
	TgID         int64  `json:"tg_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	IsAccepted   *bool  `json:"is_accepted"`
	WithSomebody *bool  `json:"with_somebody"`
	SecondGuest  string `json:"second_guest"`
	NeedTransfer string `json:"need_transfer"`
	Wishes       string `json:"wishes"`
}

func (repo *Repository) SaveAnswers(anws GuestAnswers) *logger.SlogError {
	body, err := json.Marshal(anws)
	if err != nil {
		return logger.NewSlogError(err, "error marshalling guest answers")
	}

	_, err = repo.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &repo.bucketName,
		Key:    aws.String(strconv.FormatInt(anws.TgID, 10) + ".json"),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		return logger.NewSlogError(err, "error putting object to s3")
	}

	return nil
}

func (repo *Repository) GetAnswers(tgID int64) (*GuestAnswers, *logger.SlogError) {
	resp, err := repo.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &repo.bucketName,
		Key:    aws.String(strconv.FormatInt(tgID, 10) + ".json"),
	})
	if err != nil {
		return nil, logger.NewSlogError(err, "error getting object from s3")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, logger.NewSlogError(err, "error reading object body")
	}

	var answers *GuestAnswers
	if err := json.Unmarshal(body, &answers); err != nil {
		return nil, logger.NewSlogError(err, "error unmarshalling guest answers")
	}

	return answers, nil
}

func (repo *Repository) GetAllAnswers() ([]GuestAnswers, *logger.SlogError) {
	resp, err := repo.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &repo.bucketName,
	})
	if err != nil {
		return nil, logger.NewSlogError(err, "error listing objects in s3")
	}

	var answers []GuestAnswers
	for _, obj := range resp.Contents {
		if obj.Key != nil {
			if *obj.Key == "bot_config.json" {
				continue
			}

			resp, err := repo.client.GetObject(context.TODO(), &s3.GetObjectInput{
				Bucket: &repo.bucketName,
				Key:    obj.Key,
			})
			if err != nil {
				return nil, logger.NewSlogError(err, "error getting object from s3")
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, logger.NewSlogError(err, "error reading object body")
			}

			var answer GuestAnswers
			if err := json.Unmarshal(body, &answer); err != nil {
				return nil, logger.NewSlogError(err, "error unmarshalling guest answers")
			}

			answers = append(answers, answer)
		}
	}

	return answers, nil
}
