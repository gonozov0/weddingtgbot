package lambda

import (
	"encoding/base64"
	"encoding/json"
	"log/slog"
)

func DecodeBody(rawReq []byte) ([]byte, error) {
	slog.Info("Raw request data", slog.String("req", string(rawReq)))

	req := &Request{}
	if err := json.Unmarshal(rawReq, req); err != nil {
		slog.Error("Error unmarshalling request data", slog.String("err", err.Error()))
		return nil, err
	}

	var (
		body []byte
		err  error
	)
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

	return body, nil
}
