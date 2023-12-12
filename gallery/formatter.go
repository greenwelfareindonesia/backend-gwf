package gallery

import (
	"time"
)

type GalleryFormatter struct {
	ID    int    `json:"ID"`
	Alt   string `json:"Alt"`
	GalleryImages     []string `json:"FileNames"`     
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func PostFormatterGallery(gallery Gallery) GalleryFormatter {
	formatter := GalleryFormatter{
		ID:    gallery.ID,
		Alt:   gallery.Alt,
		GalleryImages: make([]string,len(gallery.FileName)),
		CreatedAt: gallery.CreatedAt,
		UpdatedAt: gallery.UpdatedAt,
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
