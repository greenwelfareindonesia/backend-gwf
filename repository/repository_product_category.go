package repository

import (
	"context"
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryProductCategory interface {
	GetProductCategoryById(ctx context.Context, categoryId uint64) (entity.ProductCategory, error)
}

type repository_product_category struct {
	db *gorm.DB
}

func NewRepositoryProductCategory(db *gorm.DB) *repository_product_category {
	return &repository_product_category{db}
}

func (r *repository_product_category) GetProductCategoryById(ctx context.Context, categoryId uint64) (entity.ProductCategory, error) {
	productCategory := entity.ProductCategory{}

	if err := r.db.WithContext(ctx).Table("product_categories").Where("id = ?", categoryId).First(&productCategory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.ProductCategory{}, errors.New("category product not found")
		}

		return entity.ProductCategory{}, err
	}

	return productCategory, nil
}
