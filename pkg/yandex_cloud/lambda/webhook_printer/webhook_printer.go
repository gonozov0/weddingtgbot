package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

// Request represents struct that is passed to cloud function in Yandex Cloud.
type Request struct {
	HTTPMethod                      string              `json:"httpMethod"`
	Headers                         map[string]string   `json:"headers"`
	URL                             string              `json:"url"`
	Params                          map[string]string   `json:"params"`
	MultiValueParams                map[string][]string `json:"multiValueParams"`
	PathParams                      map[string]string   `json:"pathParams"`
	MultiValueHeaders               map[string][]string `json:"multiValueHeaders"`
	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`
	RequestContext                  requestContext      `json:"requestContext"`
	Body                            string              `json:"body"`
	IsBase64Encoded                 bool                `json:"isBase64Encoded"`
}

type requestContext struct {
	Identity struct {
		SourceIP  string `json:"sourceIp"`
		UserAgent string `json:"userAgent"`
	} `json:"identity"`
	HTTPMethod       string `json:"httpMethod"`
	RequestID        string `json:"requestId"`
	RequestTime      string `json:"requestTime"`
	RequestTimeEpoch int64  `json:"requestTimeEpoch"`
}

// Response represents struct that is required to be returned from cloud function in Yandex Cloud.
type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func Handler(_ context.Context, rawReq []byte) (*Response, error) {
	setupLogger()
	slog.Info("Raw request data", slog.String("req", string(rawReq)))

	req := &Request{}
	if err := json.Unmarshal(rawReq, req); err != nil {
		slog.Error("Error unmarshalling request data", slog.String("err", err.Error()))
		return nil, err
	}

	logData, err := json.Marshal(req)
	if err != nil {
		slog.Error("Error marshalling request data", slog.String("err", err.Error()))
		return nil, err
	}

	slog.Info("Marshalled request data", slog.String("req", string(logData)))

	var body []byte
	if req.IsBase64Encoded {
		body, err = base64.StdEncoding.DecodeString(req.Body)
		if err != nil {
			slog.Error("Error decoding base64 body", slog.String("err", err.Error()))
			return nil, err
		}
	} else {
		body = []byte(req.Body)
	}

	slog.Info("Decoded request body", slog.String("body", string(body)))

	return &Response{
		StatusCode: http.StatusOK,
		Body:       "ok",
	}, nil
}

func setupLogger() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}
