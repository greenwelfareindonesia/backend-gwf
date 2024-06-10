package repository

import (
	"context"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryProduct interface {
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
}

type repository_product struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repository_product {
	return &repository_product{db}
}

func (r *repository_product) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	if err := r.db.WithContext(ctx).Table("products").Save(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}
