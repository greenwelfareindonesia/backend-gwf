package dto

import (
	"errors"
	"greenwelfare/entity"
)

type CreateProductDTO struct {
	Name           string `form:"name"`
	Excerpt        string `form:"excerpt"`
	Description    string `form:"description"`
	Merk           string `form:"merk"`
	ItemWeight     string `form:"itemWeight"`
	IsActive       string `form:"isActive"`
	CategoryID     string `form:"categoryId"`
	TotalStock     uint64
	ProductImages  []CreateProductImageDTO
	ProductDetails []CreateProductDetailDTO
}

type UpdateProductDTO struct {
	Name        string `form:"name"`
	Price       uint64 `form:"price"`
	Stock       uint64 `form:"stock"`
	Description string `form:"description"`
}

type ProductResponseDTO struct {
	ID             uint64                      `json:"id"`
	Slug           string                      `json:"slug" gorm:"not null"`
	Name           string                      `json:"name" gorm:"not null"`
	CategoryID     uint64                      `json:"category_id" gorm:"foreignKey:ProductID"` //optional
	Excerpt        string                      `json:"excerpt" gorm:"not null"`
	Description    string                      `json:"description" gorm:"not null; type:text"`
	Merk           string                      `json:"merk" gorm:"not null"`
	TotalStock     uint64                      `json:"total_stock" gorm:"not null"`
	TotalSales     uint64                      `json:"total_sales" gorm:"default:0"`
	ItemWeight     float64                     `json:"item_weight" gorm:"not null"`
	IsActive       bool                        `json:"is_active" gorm:"default:true"`
	ProductImages  []*ProductImageResponseDTO  `json:"product_images"`
	ProductDetails []*ProductDetailResponseDTO `json:"product_details"`
	entity.DefaultColumn
}

func (p *CreateProductDTO) Validate() error {
	if p.Name == "" {
		return errors.New("field name is required")
	}

	if p.Excerpt == "" {
		return errors.New("field except is required")
	}

	if p.Description == "" {
		return errors.New("field description is requried")
	}

	if p.Merk == "" {
		return errors.New("field merk is requried")
	}

	if p.ItemWeight == "" {
		return errors.New("field itemWeight is requried")
	}

	if p.IsActive == "" {
		return errors.New("field isActive is requried")
	}
	return nil
}
