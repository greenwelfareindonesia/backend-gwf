package entity

type ShoppingCart struct {
	ID              uint64        `json:"id" gorm:"primaryKey"`
	UserID          uint64        `json:"user_id" gorm:"not null; foreignKey:UserID"`
	ProductID       uint64        `json:"product_id" gorm:"not null; foreignKey:ProductID"`
	ProductDetailID uint64        `json:"product_detail_id" gorm:"not null; foreignKey:ProductDetailID"`
	Qty             uint64        `json:"qty" gorm:"not null"`
	Price           uint64        `json:"price" gorm:"not null"`
	TotalPrice      uint64        `json:"total_price" gorm:"not null"`
	Product         Product       `json:"product"`
	ProductDetail   ProductDetail `json:"product_detail"`
	User            User          `json:"user"`
	DefaultColumn
}
