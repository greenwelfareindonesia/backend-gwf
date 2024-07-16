package dto

import "greenwelfare/entity"

type CreateProductImageDTO struct {
	ImageUrl string `form:"image"`
}

type ProductImageResponseDTO struct {
	ID        uint64         `json:"id"`
	ProductID uint64         `json:"product_id"`
	ImageUrl  string         `json:"image_url"`
	Product   *entity.Product `json:"product,omitempty"`
	entity.DefaultColumn
}
