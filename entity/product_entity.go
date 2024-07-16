package entity

type Product struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Slug        string `json:"slug" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Price       uint64 `json:"price" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	ImageUrl    string `json:"image_url" gorm:"not null"`
	Stock       uint64 `json:"stock" gorm:"not null"`
	DefaultColumn
}
