package formatter

import (
	"greenwelfare/entity"
	"time"
)

type CreateArtikelFormatter struct {
	ID             int `json:"ID"`
	Slug           string `json:"slug"`
	FullName       string    `json:"fullName"`
	Email string `json:"email"`
	Topic string `json:"topic"`
	ArticleMessage string    `json:"articleMessage"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func PostArticleFormat(artikel *entity.Artikel) CreateArtikelFormatter {
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

func UpdatedArticleFormat(artikel *entity.Artikel) CreateArtikelFormatter {
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


func FormatterGetArtikel(artikel []*entity.Artikel) []CreateArtikelFormatter {
	artikelGetFormatter := []CreateArtikelFormatter{}

	for _, artikels := range artikel {
		artikelsFormatter := PostArticleFormat(artikels)
		artikelGetFormatter = append(artikelGetFormatter, artikelsFormatter)
	}

	return artikelGetFormatter
}