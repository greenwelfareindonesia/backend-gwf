package dto

import (
	"errors"
	"greenwelfare/entity"
)

type CreateProductDTO struct {
	Name        string `form:"name"`
	Price       uint64 `form:"price"`
	Stock       uint64 `form:"stock"`
	Description string `form:"description"`
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

	if p.Price < 1 || p.Price == 0 {
		return errors.New("field price min 1")
	}

	if p.Stock < 1 || p.Stock == 0 {
		return errors.New("field stock min 1")
	}

	if p.Description == "" {
		return errors.New("field description is requried")
	}

	return nil
}
