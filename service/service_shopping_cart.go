package service

import (
	"context"
	// "errors"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"
)

type ServiceShoppingCart interface {
	CreateShoppingCart(ctx context.Context, shoppingCart dto.CreateShoppingCartDTO) (dto.ShoppingCartResponseDTO, error)
	GetShoppingCarts(ctx context.Context, userId uint64) ([]dto.ShoppingCartResponseDTO, error)
	GetShoppingCartById(ctx context.Context, userId uint64, cardId uint64) (dto.ShoppingCartResponseDTO, error)
	UpdateShoppingCartById(ctx context.Context, updateShoppingCart dto.UpdateShoppingCartDTO) (dto.ShoppingCartResponseDTO, error) // update qty and total price
	DeleteShoppingCartById(ctx context.Context, userId uint64, cardId uint64) error
	GetStatisticCarts(ctx context.Context, userId uint64) (dto.ShoppingCartStatisticResponseDTO, error)
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
	// existingProduct, errRepoProduct := s.repoProduct.GetProductById(ctx, newShoppingCart.ProductID)
	// if errRepoProduct != nil {
	// 	return dto.ShoppingCartResponseDTO{}, errRepoProduct
	// }

	// VALIDATE QTY PRODUCT WITH QTY SHOPPING CART
	// if existingProduct.Stock < shoppingCart.Qty {
	// 	return dto.ShoppingCartResponseDTO{}, errors.New("out of stock")
	// }

	// TOTAL PRODUCT PRICE WITH QTY SHOPPING CART
	// totalPrice := existingProduct.Price * shoppingCart.Qty
	// newShoppingCart.TotalPrice = totalPrice

	// SAVE DATA TO DB
	res, errRepoShoppingCart := s.repoShoppingCart.CreateShoppingCart(ctx, newShoppingCart)
	if errRepoShoppingCart != nil {
		return dto.ShoppingCartResponseDTO{}, errRepoShoppingCart
	}
	return parsingShoppingCartResponseDTO(res), nil
}

func (s *service_shopping_cart) GetShoppingCarts(ctx context.Context, userId uint64) ([]dto.ShoppingCartResponseDTO, error) {
	res, errRepo := s.repoShoppingCart.GetShoppingCarts(ctx, userId)
	if errRepo != nil {
		return []dto.ShoppingCartResponseDTO{}, errRepo
	}

	resposeDTO := []dto.ShoppingCartResponseDTO{}
	for _, v := range res {
		parsingRes := parsingShoppingCartResponseDTO(v)
		resposeDTO = append(resposeDTO, parsingRes)
	}

	return resposeDTO, nil
}

func (s *service_shopping_cart) GetShoppingCartById(ctx context.Context, userId uint64, cardId uint64) (dto.ShoppingCartResponseDTO, error) {
	shoppingCart, errRepo := s.repoShoppingCart.GetShoppingCartById(ctx, userId, cardId)
	if errRepo != nil {
		return dto.ShoppingCartResponseDTO{}, errRepo
	}

	return parsingShoppingCartResponseDTO(shoppingCart), nil
}

func (s *service_shopping_cart) UpdateShoppingCartById(ctx context.Context, updateShoppingCart dto.UpdateShoppingCartDTO) (dto.ShoppingCartResponseDTO, error) {
	// GET EXISTING SHOPPING CART
	shoppingCart, errRepo := s.repoShoppingCart.GetShoppingCartById(ctx, updateShoppingCart.UserID, updateShoppingCart.ID)
	if errRepo != nil {
		return dto.ShoppingCartResponseDTO{}, errRepo
	}

	// GET DATA NEW PRODUCT FROM UPDATED SHOPPING CART
	// product, errRepoProduct := s.repoProduct.GetProductById(ctx, shoppingCart.ProductID)
	// if errRepoProduct != nil {
	// 	return dto.ShoppingCartResponseDTO{}, errRepoProduct
	// }

	// VALIDATE QTY PRODUCT WITH QTY SHOPPING CART
	// if product.Stock < shoppingCart.Qty {
	// 	return dto.ShoppingCartResponseDTO{}, errors.New("out of stock")
	// }

	// UPDATE SHOPPING CART
	if updateShoppingCart.Status == "increment" {
		shoppingCart.Qty = shoppingCart.Qty + 1
	}

	if updateShoppingCart.Status == "decrement" {
		shoppingCart.Qty = shoppingCart.Qty - 1
	}

	// shoppingCart.TotalPrice = (shoppingCart.Qty * product.Price)

	updatedData, errRepo := s.repoShoppingCart.UpdateShoppingCartById(ctx, shoppingCart)
	if errRepo != nil {
		return dto.ShoppingCartResponseDTO{}, errRepo
	}

	return parsingShoppingCartResponseDTO(updatedData), nil
}

func (s *service_shopping_cart) DeleteShoppingCartById(ctx context.Context, userId uint64, cardId uint64) error {
	shoppingCart, errRepo := s.repoShoppingCart.GetShoppingCartById(ctx, userId, cardId)
	if errRepo != nil {
		return errRepo
	}

	errRepo = s.repoShoppingCart.DeleteShoppingCartById(ctx, shoppingCart.ID)
	if errRepo != nil {
		return errRepo
	}

	return nil
}

func (s *service_shopping_cart) GetStatisticCarts(ctx context.Context, userId uint64) (dto.ShoppingCartStatisticResponseDTO, error) {
	res, errRepo := s.repoShoppingCart.GetShoppingCarts(ctx, userId)
	if errRepo != nil {
		return dto.ShoppingCartStatisticResponseDTO{}, errRepo
	}

	resposeDTO := dto.ShoppingCartStatisticResponseDTO{}
	total_item := 0
	total_product := len(res)

	for _, v := range res {
		total_item += int(v.Qty)
	}
	
	resposeDTO.TotalItem = uint64(total_item)
	resposeDTO.TotalProduct = uint64(total_product)

	return resposeDTO, nil
}

func parsingShoppingCartResponseDTO(shoppingCart entity.ShoppingCart) dto.ShoppingCartResponseDTO {
	response := dto.ShoppingCartResponseDTO{
		ID:            shoppingCart.ID,
		ProductID:     shoppingCart.ProductID,
		Qty:           shoppingCart.Qty,
		TotalPrice:    shoppingCart.TotalPrice,
		DefaultColumn: shoppingCart.DefaultColumn,
	}

	if shoppingCart.User.ID != 0 {
		userResponse := dto.UserShoppingCartResponse{
			ID:       uint64(shoppingCart.User.ID),
			Username: shoppingCart.User.Username,
		}
		response.User = &userResponse
	}

	if shoppingCart.Product.ID != 0 {
		productResponse := dto.ProductShoppingCartResponse{
			ID:          shoppingCart.Product.ID,
			Name:        shoppingCart.Product.Name,
			// Price:       shoppingCart.Product.Price,
			// ImageUrl:    shoppingCart.Product.ImageUrl,
			Description: shoppingCart.Product.Description,
		}
		response.Product = &productResponse
	}

	return response
}
