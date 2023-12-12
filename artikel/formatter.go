package artikel

import "time"

type CreateArtikelFormatter struct {
	ID             int `json:"ID"`
	Slug           string `json:"Slug"`
	FullName       string    `json:"FullName"`
	Email string `json:"Email"`
	Topic string `json:"Topic"`
	ArticleMessage string    `json:"ArticleMessage"`
	CreatedAt      time.Time `json:"CreatedAt"`
	UpdatedAt      time.Time `json:"UpdatedAt"`
}

func PostArticleFormat(artikel Artikel) CreateArtikelFormatter {
	formatter := CreateArtikelFormatter{
		ID:        artikel.ID,
		Slug: artikel.Slug,
		FullName: artikel.FullName,
		Email: artikel.Email,
		Topic: artikel.Topic,
		ArticleMessage:   artikel.ArticleMessage,
		CreatedAt: artikel.CreatedAt,
	}
	return formatter
}

func UpdatedArticleFormat(artikel Artikel) CreateArtikelFormatter {
	formatter := CreateArtikelFormatter{
		ID:        artikel.ID,
		Slug: artikel.Slug,
		FullName: artikel.FullName,
		Email: artikel.Email,
		Topic: artikel.Topic,
		ArticleMessage:   artikel.ArticleMessage,
		UpdatedAt: artikel.UpdatedAt,
	}
	return formatter
}


func FormatterGetArtikel(artikel []Artikel) []CreateArtikelFormatter {
	artikelGetFormatter := []CreateArtikelFormatter{}

	for _, artikels := range artikel {
		artikelsFormatter := PostArticleFormat(artikels)
		artikelGetFormatter = append(artikelGetFormatter, artikelsFormatter)
	}

	return artikelGetFormatter
}