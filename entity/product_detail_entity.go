package entity

type ProductDetail struct {
	ID        uint64  `json:"id" gorm:"primaryKey"`
	ProductID uint64  `json:"product_id" gorm:"not null; foreignKey:ProductID"`
	Size      string  `json:"size" gorm:"not null"`
	Price     uint64   `json:"price" gorm:"not null; nim:1"`
	Stock     uint64  `json:"stock" gorm:"not null"`
	Product   Product `json:"product"`
	DefaultColumn
}
