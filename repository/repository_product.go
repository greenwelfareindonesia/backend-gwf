package repository

import (
	"context"
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryProduct interface {
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	GetProductById(ctx context.Context, id uint64) (entity.Product, error)
	ReadProductBySlug(ctx context.Context, slug string) (entity.Product, error)
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

func (r *repository_product) GetProductById(ctx context.Context, id uint64) (entity.Product, error) {
	product := entity.Product{}

	if err := r.db.WithContext(ctx).Table("products").Where("id = ?", id).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Product{}, errors.New("product not found")
		}

		return entity.Product{}, err
	}

	return product, nil
}

func (r *repository_product) ReadProductBySlug(ctx context.Context, slug string) (entity.Product, error) {
	var product entity.Product

	if err := r.db.WithContext(ctx).Table("products").Where("slug = ?", slug).First(&product).Error; err != nil {
		return entity.Product{}, err
	}

	return product, nil
}
