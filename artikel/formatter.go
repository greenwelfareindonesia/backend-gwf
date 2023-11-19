package artikel

import "time"

type CreateArtikelFormatter struct {
	ID             int `json:"id"`
	FullName       string    `json:"full_name"`
	Email string `json:"email"`
	Topic string `json:"topic"`
	ArtikelMessage string    `json:"message"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FormatterArtikel (artikel Artikel) CreateArtikelFormatter {
	formatter := CreateArtikelFormatter{
		ID:        artikel.ID,
		FullName: artikel.FullName,
		Email: artikel.Email,
		Topic: artikel.Topic,
		ArtikelMessage:   artikel.ArtikelMessage,
		CreatedAt: artikel.CreatedAt,
	}
	return formatter
}

func FormatterGetArtikel(artikel []Artikel) []CreateArtikelFormatter {
	artikelGetFormatter := []CreateArtikelFormatter{}

	for _, artikels := range artikel {
		artikelsFormatter := FormatterArtikel(artikels)
		artikelGetFormatter = append(artikelGetFormatter, artikelsFormatter)
	}

	return artikelGetFormatter
}