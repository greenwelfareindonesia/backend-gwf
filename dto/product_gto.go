package dto

import (
	"errors"
	"greenwelfare/entity"
)

type CreateProductDTO struct {
	Name        string `form:"name" binding:"required"`
	Price       uint64 `form:"price" binding:"required"`
	Stock       uint64 `form:"stock" binding:"required"`
	Description string `form:"description" binding:"required"`
}

type ProductResponseDTO struct {
	ID          uint64   `json:"ID"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Price       uint64 `json:"price"`
	Description string `json:"description"`
	Stock       uint64 `json:"stock"`
	entity.DefaultColumn
}

func (p *CreateProductDTO) Validate() error {
	if p.Stock < 1 {
		return errors.New("Product Stock min 1")
	}
	return nil
}
