package entity

type ProductImage struct {
	ID        uint64  `json:"id" gorm:"primaryKey"`
	ProductID uint64  `json:"product_id" gorm:"not null; foreignKey:ProductID"`
	ImageUrl  string  `json:"image_url" gorm:"not null"`
	Product   Product `json:"product"`
	DefaultColumn
}
