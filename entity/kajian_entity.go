package entity

type Kajian struct {
	ID          int
	Slug        string
	Title       string
	Description string
	Images      []KajianImage `gorm:"foreignKey:KajianID"`
	DefaultColumn
}

type KajianImage struct {
	ID        int
	FileName  string
	KajianID  int
	DefaultColumn
}
