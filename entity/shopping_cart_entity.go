package entity

type ShoppingCart struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	ProductID   uint64  `json:"product_id" gorm:"not null; foreignKey:ProductID"`
	UserID      uint64  `json:"user_id" gorm:"not null; foreignKey:UserID"`
	Qty         uint64  `json:"qty" gorm:"not null"`
	TotalPrice uint64  `json:"total_price" gorm:"not null"`
	Product     Product `json:"product"`
	User        User    `json:"user"`
	DefaultColumn
}
