package entity

type Ecopedia struct {
	ID          int
	Slug        string
	Title       string
	SubTitle    string
	Description string
	SrcFile     string
	Reference   string
	FileName    []EcopediaImage `gorm:"foreignKey:EcopediaID"`
	DefaultColumn
}

type EcopediaImage struct {
	ID         int `gorm:"primaryKey"`
	FileName   string
	EcopediaID int
	DefaultColumn
}
