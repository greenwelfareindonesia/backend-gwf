package gallery

import (
	"time"
)

type Gallery struct {
	ID    int
	Slug string
	Alt   string
	FileName []GalleryImages `gorm:"foreignKey:GalleryID"`
	ActionUserGallery []ActionUsersGallery `gorm:"foreignKey:GalleryID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GalleryImages struct {
	ID int
	FileName string
	GalleryID int
	CreatedAt time.Time
	UpdatedAt time.Time
 }

type ActionUsersGallery struct {
	ID int
	Like bool
	GalleryID int
	CreatedAt time.Time
	UpdatedAt time.Time
}
