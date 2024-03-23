package s3

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gonozov0/weddingtgbot/pkg/logger"
)

type Repository struct {
	client     *s3.Client
	bucketName string
}

func NewRepository() (*Repository, *logger.SlogError) {
	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
			if service == s3.ServiceID && region == "ru-central1" {
				return aws.Endpoint{
					PartitionID:   "yc",
					URL:           "https://storage.yandexcloud.net",
					SigningRegion: "ru-central1",
				}, nil
			}
			return aws.Endpoint{}, errors.New("unknown endpoint requested")
		},
	)

	accessKeyID := os.Getenv("YC_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("YC_SECRET_ACCESS_KEY")
	bucketName := os.Getenv("YC_BUCKET_NAME")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("ru-central1"),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		return nil, logger.NewSlogError(err, "error loading S3 config")
	}

	client := s3.NewFromConfig(cfg)

	return &Repository{
		client:     client,
		bucketName: bucketName,
	}, nil
}
