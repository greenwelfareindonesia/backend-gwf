package dto

import (
	"errors"
	"greenwelfare/entity"
)

type CreateProductDetailDTO struct {
	Size  string `form:"size"`
	Price uint64 `form:"price"`
	Stock uint64 `form:"stock"`
}

func (pd *CreateProductDetailDTO) Validate() error {
	if pd.Size == "" {
		return errors.New("field size is required")
	}

	if pd.Price < 1 {
		return errors.New("field price min 1")
	}

	if pd.Stock < 1 {
		return errors.New("field stock min 1")
	}

	return nil
}

type ProductDetailResponseDTO struct {
	ID      uint64         `json:"id"`
	Size    string         `json:"size"`
	Price   uint64         `json:"price"`
	Stock   uint64         `json:"stock"`
	Product *entity.Product `json:"product,omitempty"`
	entity.DefaultColumn
}
