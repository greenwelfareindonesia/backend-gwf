package entity

type Banner struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"not null"`
	ImageUrl string `json:"image_url" gorm:"not null"`
	DefaultColumn
}
