package gallery

import (
	"time"
)

type GalleryFormatter struct {
	ID    int    `json:"ID"`
	Alt   string `json:"Alt"`
	Slug string `json:"Slug"`
	GalleryImages     []string `json:"FileNames"`     
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func PostFormatterGallery(gallery Gallery) GalleryFormatter {
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

func FormatterGetAllGallery (gallery []Gallery) []GalleryFormatter {
	newGalleryGetFormatter := []GalleryFormatter{}

	for _, newGallery := range gallery {
		newGalleryFormatter := PostFormatterGallery(newGallery)
		newGalleryGetFormatter = append(newGalleryGetFormatter, newGalleryFormatter)
	}

	return newGalleryGetFormatter
}
