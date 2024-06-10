package repository

import (
	"context"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryShoppingCart interface {
	CreateShoppingCart(ctx context.Context, shoppingCart entity.ShoppingCart) (entity.ShoppingCart, error)
	GetShoppingCarts(ctx context.Context, userId uint64) ([]entity.ShoppingCart, error)
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

func (r *repository_shopping_cart) GetShoppingCarts(ctx context.Context, userId uint64) ([]entity.ShoppingCart, error) {
	shoppingCarts := []entity.ShoppingCart{}

	query := r.db.WithContext(ctx).Table("shopping_carts").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "price", "image_url", "description")
	})

	// optional filter by sessoin user_id for customer / filer user_id by admin
	if userId != 0 {
		query = query.Where("user_id = ?", userId)
	}

	if err := query.Find(&shoppingCarts).Error; err != nil {
		return []entity.ShoppingCart{}, err
	}

	return shoppingCarts, nil
}
