package main

import (
	"github.com/gonozov0/weddingtgbot/internal/repository/googledoc"
)

func main() {
	gglRepo, err := googledoc.NewRepository(
		"creds/avid-sphere-413708-c20f667d42e2.json",
		"1csSzXWCeDkDZirCwae4ozG0i0PnMmT42dttnmq8iDMo",
	)
	if err != nil {
		panic(err)
	}
	f := false

	err = gglRepo.InsertAnswers([]googledoc.AnswerDTO{
		{
			FirstName:    "Иван",
			LastName:     "Иванов",
			IsAccepted:   &f,
			WithSomebody: nil,
			SecondGuest:  "е",
			NeedTransfer: "нет",
			Wishes:       "пожалуйста, не играйте на баяне",
		},
		{
			FirstName:    "g",
			LastName:     "Петров",
			IsAccepted:   &f,
			WithSomebody: nil,
			SecondGuest:  "нет",
			NeedTransfer: "нет",
			Wishes:       "пожалуйста, не играйте на баяне",
		},
	})
	if err != nil {
		panic(err)
	}
}
