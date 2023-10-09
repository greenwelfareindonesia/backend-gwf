package gallery

import (
	"time"
)

type GalleryFormatter struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Alt   string `json:"alt"`
	//Likes int `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatterGallery(gallery Gallery) GalleryFormatter {
	formatter := GalleryFormatter{
		ID:    gallery.ID,
		Image: gallery.Image,
		Alt:   gallery.Alt,
		//Likes: gallery.Likes,
		CreatedAt: gallery.CreatedAt,
		UpdatedAt: gallery.UpdatedAt,
	}
	return formatter
}
