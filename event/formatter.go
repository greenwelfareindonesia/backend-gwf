package event

import "time"

type CreateEventFormatter struct {
	ID             int `json:"id"`
	Judul       string    `json:"judul"`
	EventMessage string    `json:"message"`
	FileName string    `json:"file_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FormatterEvent (artikel Event) CreateEventFormatter {
	formatter := CreateEventFormatter{
		ID:        artikel.ID,
		Judul: artikel.Judul,
		EventMessage:   artikel.EventMessage,
		FileName: artikel.FileName,
		CreatedAt: artikel.CreatedAt,
		UpdatedAt: artikel.UpdatedAt,
	}
	return formatter
}

func FormatterGetArtikel(artikel []Event) []CreateEventFormatter {
	artikelGetFormatter := []CreateEventFormatter{}

	for _, artikels := range artikel {
		artikelsFormatter := FormatterEvent(artikels)
		artikelGetFormatter = append(artikelGetFormatter, artikelsFormatter)
	}

	return artikelGetFormatter
}