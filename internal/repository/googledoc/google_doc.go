package googledoc

import (
	"context"

	"github.com/gonozov0/weddingtgbot/pkg/logger"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type AnswerDTO struct {
	FirstName    string
	LastName     string
	IsAccepted   *bool
	WithSomebody *bool
	SecondGuest  string
	NeedTransfer string
	Wishes       string
}

type Repository struct {
	service       *sheets.Service
	spreadsheetID string
}

func NewRepository(serviceAccountPath, spreadsheetID string) (*Repository, *logger.SlogError) {
	srv, err := sheets.NewService(context.Background(), option.WithCredentialsFile(serviceAccountPath))
	if err != nil {
		return nil, logger.NewSlogError(err, "error creating google sheets service")
	}
	return &Repository{
		service:       srv,
		spreadsheetID: spreadsheetID,
	}, nil
}

func (repo *Repository) InsertAnswer(dto AnswerDTO) *logger.SlogError {
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{
			{
				dto.FirstName,
				dto.LastName,
				dto.IsAccepted,
				dto.WithSomebody,
				dto.SecondGuest,
				dto.NeedTransfer,
				dto.Wishes,
			},
		},
	}

	_, err := repo.service.Spreadsheets.Values.Append(repo.spreadsheetID, "A1", valueRange).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		return logger.NewSlogError(err, "error appending values to google sheets")
	}

	return nil
}

func (repo *Repository) InsertAnswers(dtos []AnswerDTO) *logger.SlogError {
	var values [][]interface{}
	for _, dto := range dtos {
		values = append(values, []interface{}{
			dto.FirstName,
			dto.LastName,
			dto.IsAccepted,
			dto.WithSomebody,
			dto.SecondGuest,
			dto.NeedTransfer,
			dto.Wishes,
		})
	}
	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := repo.service.Spreadsheets.Values.Update(repo.spreadsheetID, "A2", valueRange).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		return logger.NewSlogError(err, "error appending values to google sheets")
	}

	return nil
}
