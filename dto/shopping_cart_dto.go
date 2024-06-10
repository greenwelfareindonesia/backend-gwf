package dto

import (
	"errors"
	"greenwelfare/entity"
)

type CreateShoppingCartDTO struct {
	ProductID uint64 `json:"product_id" form:"product_id"`
	Qty       uint64 `json:"qty" form:"qty"`
	UserID    uint64 `json:"user_id" form:"user_id"`
}

type ShoppingCartResponseDTO struct {
	ID         uint64                       `json:"id"`
	ProductID  uint64                       `json:"product_id"`
	Qty        uint64                       `json:"qty"`
	TotalPrice uint64                       `json:"total_price"`
	User       *UserShoppingCartResponse    `json:"user,omitempty"`
	Product    *ProductShoppingCartResponse `json:"product,omitempty"`
	entity.DefaultColumn
}

type UserShoppingCartResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

type ProductShoppingCartResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Price       uint64 `json:"price"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

func (p *CreateShoppingCartDTO) Validate() error {
	if p.Qty < 1 || p.Qty == 0 {
		return errors.New("field quantity min 1")
	}

	if p.ProductID == 0 {
		return errors.New("field product_id is required")
	}

	return nil
}
