package entity

type ProductCategory struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	DefaultColumn
}
