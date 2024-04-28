package formatter

import (
	"greenwelfare/entity"
	"time"
)

type GalleryFormatter struct {
	ID    int    `json:"ID"`
	Alt   string `json:"alt"`
	Slug string `json:"slug"`
	GalleryImages     []string `json:"fileNames"`     
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func PostFormatterGallery(gallery *entity.Gallery) GalleryFormatter {
	formatter := GalleryFormatter{
		ID:    gallery.ID,
		Slug: gallery.Slug,
		Alt:   gallery.Alt,
		CreatedAt: gallery.CreatedAt,
		UpdatedAt: gallery.UpdatedAt,
	}

	for _, fileName := range gallery.FileName {
		formatter.GalleryImages = append(formatter.GalleryImages, fileName.FileName)
	}

	return formatter
}

func FormatterGetAllGallery (gallery []*entity.Gallery) []GalleryFormatter {
	newGalleryGetFormatter := []GalleryFormatter{}

	for _, newGallery := range gallery {
		newGalleryFormatter := PostFormatterGallery(newGallery)
		newGalleryGetFormatter = append(newGalleryGetFormatter, newGalleryFormatter)
	}

	return newGalleryGetFormatter
}
