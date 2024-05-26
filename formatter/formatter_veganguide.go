package formatter

import (
	"greenwelfare/entity"
	"time"
)

type VeganguideFormatter struct {
	ID        int    `json:"ID"`
	Slug      string `json:"slug"`
	Judul     string `json:"title"`
	Deskripsi string `json:"description"`
	Body      string `json:"body"`
	Gambar    string `json:"fileName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


func PostFormatterVeganguide(vegan *entity.Veganguide) VeganguideFormatter {
	formatter := VeganguideFormatter{
		ID:        vegan.ID,
		Slug:      vegan.Slug,
		Judul:     vegan.Judul,
		Deskripsi:     vegan.Deskripsi,
		Body:      vegan.Body,
		Gambar:      vegan.Gambar,
		CreatedAt: vegan.CreatedAt,
	}
	return formatter
}

func UpdatedFormatterVeganguide(vegan *entity.Veganguide) VeganguideFormatter {
	formatter := VeganguideFormatter{
		ID:        vegan.ID,
		Slug:      vegan.Slug,
		Judul:     vegan.Judul,
		Deskripsi:     vegan.Deskripsi,
		Body:      vegan.Body,
		Gambar:      vegan.Gambar,
		UpdatedAt: vegan.UpdatedAt,
	}
	return formatter
}