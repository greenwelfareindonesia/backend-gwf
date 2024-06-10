package repository

import (
	"context"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryShoppingCart interface {
	CreateShoppingCart(ctx context.Context, shoppingCart entity.ShoppingCart) (entity.ShoppingCart, error)
}

type repository_shopping_cart struct {
	db *gorm.DB
}

func NewRepositoryShoppingCart(db *gorm.DB) *repository_shopping_cart {
	return &repository_shopping_cart{db}
}

func (r *repository_shopping_cart) CreateShoppingCart(ctx context.Context, shoppingCart entity.ShoppingCart) (entity.ShoppingCart, error) {
	if err := r.db.WithContext(ctx).Table("shopping_carts").Save(&shoppingCart).Error; err != nil {
		return entity.ShoppingCart{}, err
	}
	return shoppingCart, nil
}
