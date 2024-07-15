package repository

import (
	"context"
	// "errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryProductImage interface {
	CreateProductImage(ctx context.Context, product entity.Product) (entity.Product, error)
	// GetProductImageById(ctx context.Context, id uint64) (entity.Product, error)
	// ReadProductImageBySlug(ctx context.Context, slug string) (entity.Product, error)
	// UpdateProductImage(ctx context.Context, product *entity.Product) (entity.Product, error)
	// DeleteProductImage(ctx context.Context, product *entity.Product) (entity.Product, error)
}

type repository_product_image struct {
	db *gorm.DB
}

func NewRepositoryProductImage(db *gorm.DB) *repository_product_image {
	return &repository_product_image{db}
}