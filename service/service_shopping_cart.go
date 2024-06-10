package service

import (
	"context"
	"errors"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"
)

type ServiceShoppingCart interface {
	CreateShoppingCart(ctx context.Context, shoppingCart dto.CreateShoppingCartDTO) (dto.ShoppingCartResponseDTO, error)
}

type service_shopping_cart struct {
	repoShoppingCart repository.RepositoryShoppingCart
	repoProduct      repository.RepositoryProduct
}

func NewServiceShoppingCart(repoShoppingCart repository.RepositoryShoppingCart, repoProduct repository.RepositoryProduct) *service_shopping_cart {
	return &service_shopping_cart{repoShoppingCart, repoProduct}
}

func (s *service_shopping_cart) CreateShoppingCart(ctx context.Context, shoppingCart dto.CreateShoppingCartDTO) (dto.ShoppingCartResponseDTO, error) {
	newShoppingCart := entity.ShoppingCart{
		ProductID: shoppingCart.ProductID,
		Qty:       shoppingCart.Qty,
		UserID:    shoppingCart.UserID,
	}

	// GET DATA PRODUCT
	existingProduct, errRepoProduct := s.repoProduct.GetProductById(ctx, newShoppingCart.ProductID)
	if errRepoProduct != nil {
		return dto.ShoppingCartResponseDTO{}, errRepoProduct
	}

	// VALIDATE QTY PRODUCT WITH QTY SHOPPING CART
	if existingProduct.Stock < shoppingCart.Qty {
		return dto.ShoppingCartResponseDTO{}, errors.New("out of stock")
	}

	// TOTAL PRODUCT PRICE WITH QTY SHOPPING CART
	totalPrice := existingProduct.Price * shoppingCart.Qty
	newShoppingCart.TotalPrice = totalPrice

	// SAVE DATA TO DB
	res, errRepoShoppingCart := s.repoShoppingCart.CreateShoppingCart(ctx, newShoppingCart)
	if errRepoShoppingCart != nil {
		return dto.ShoppingCartResponseDTO{}, errRepoShoppingCart
	}
	return parsingShoppingCartResponseDTO(res), nil
}

func parsingShoppingCartResponseDTO(shoppingCart entity.ShoppingCart) dto.ShoppingCartResponseDTO {
	response := dto.ShoppingCartResponseDTO{
		ID:            shoppingCart.ID,
		ProductID:     shoppingCart.ProductID,
		Qty:           shoppingCart.Qty,
		TotalPrice:    shoppingCart.TotalPrice,
		DefaultColumn: shoppingCart.DefaultColumn,
	}
	return response
}
