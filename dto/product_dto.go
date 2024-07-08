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
	ItemWeight     string `form:"item_weight"`
	IsActive       string `form:"is_active"`
	CategoryID     string `form:"category_id"`
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
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Price       uint64 `json:"price"`
	Description string `json:"description"`
	Stock       uint64 `json:"stock"`
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
		return errors.New("field item_weight is requried")
	}

	if p.IsActive == "" {
		return errors.New("field is_active is requried")
	}
	return nil
}
