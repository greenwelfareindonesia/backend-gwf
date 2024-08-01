package entity

type Product struct {
	ID             uint64          `json:"id" gorm:"primaryKey"`
	Slug           string          `json:"slug" gorm:"not null"`
	Name           string          `json:"name" gorm:"not null"`
	CategoryID     uint64          `json:"category_id" gorm:"foreignKey:ProductID"` //optional
	Excerpt        string          `json:"excerpt" gorm:"not null"`
	Description    string          `json:"description" gorm:"not null; type:text"`
	Merk           string          `json:"merk" gorm:"not null"`
	TotalStock     uint64          `json:"total_stock" gorm:"not null"`
	TotalSales     uint64          `json:"total_sales" gorm:"default:0"`
	ItemWeight     float64         `json:"item_weight" gorm:"not null"`
	IsActive       bool            `json:"is_active" gorm:"default:true"`
	ProductImages  []ProductImage  `json:"product_images" gorm:"foreignKey:ProductID"`
	ProductDetails []ProductDetail `json:"product_details" gorm:"foreignKey:ProductID"`
	DefaultColumn
}
