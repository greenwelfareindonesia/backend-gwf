package entity

type Gallery struct {
	ID       int
	Slug     string
	Alt      string
	FileName []GalleryImages `gorm:"foreignKey:GalleryID"`
	DefaultColumn
}

type GalleryImages struct {
	ID        int
	FileName  string
	GalleryID int
	DefaultColumn
}
