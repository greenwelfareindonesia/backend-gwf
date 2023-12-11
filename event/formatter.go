package event

import "time"

type CreateEventFormatter struct {
	ID           int       `json:"ID"`
	Judul        string    `json:"Judul"`
	EventMessage string    `json:"Message"`
	FileName     string    `json:"FileName"`
	CreatedAt    time.Time `json:"CreatedAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
}

func PostFormatterEvent(artikel Event) CreateEventFormatter {
	formatter := CreateEventFormatter{
		ID:           artikel.ID,
		Judul:        artikel.Judul,
		EventMessage: artikel.EventMessage,
		FileName:     artikel.FileName,
		CreatedAt:    artikel.CreatedAt,
		// UpdatedAt:    artikel.UpdatedAt,
	}
	return formatter
}

func UpdatedFormatterEvent(artikel Event) CreateEventFormatter {
	formatter := CreateEventFormatter{
		ID:           artikel.ID,
		Judul:        artikel.Judul,
		EventMessage: artikel.EventMessage,
		FileName:     artikel.FileName,
		UpdatedAt:    artikel.UpdatedAt,
	}
	return formatter
}

func FormatterGetArtikel(artikel []Event) []CreateEventFormatter {
	artikelGetFormatter := []CreateEventFormatter{}

	for _, artikels := range artikel {
		artikelsFormatter := PostFormatterEvent(artikels)
		artikelGetFormatter = append(artikelGetFormatter, artikelsFormatter)
	}

	return artikelGetFormatter
}
