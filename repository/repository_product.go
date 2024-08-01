package repository

import (
	"context"
	"errors"
	"greenwelfare/entity"

	"gorm.io/gorm"
)

type RepositoryProduct interface {
	CreateProduct(ctx context.Context, product entity.Product, productImages []entity.ProductImage, productDetails []entity.ProductDetail) (entity.Product, error)
	GetProducts(ctx context.Context, limit int, offset int) ([]entity.Product, error)
	GetProductById(ctx context.Context, id uint64) (entity.Product, error)
	ReadProductBySlug(ctx context.Context, slug string) (entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) (entity.Product, error)
	DeleteProduct(ctx context.Context, product *entity.Product) (entity.Product, error)
}

type repository_product struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repository_product {
	return &repository_product{db}
}

func (r *repository_product) CreateProduct(ctx context.Context, product entity.Product, productImages []entity.ProductImage, productDetails []entity.ProductDetail) (entity.Product, error) {
	// USING DB TRANSACTIONAL
	tx := r.db.Begin()

	// CREATE PRODUCT
	if err := r.db.WithContext(ctx).Table("products").Save(&product).Error; err != nil {
		// ROLLBACK ERROR WITH DELETE IMAGEKIT
		tx.Rollback()
		return entity.Product{}, err
	}

	// CREATE PRODUCT IMAGE
	for _, productImage := range productImages {
		productImage.ProductID = product.ID
		if err := r.db.WithContext(ctx).Table("product_images").Save(&productImage).Error; err != nil {
			// ROLLBACK ERROR WITH DELETE IMAGEKIT
			tx.Rollback()
			return entity.Product{}, err
		}
	}

	// CREATE PRODUCT DETAIL
	for _, productDetail := range productDetails {
		productDetail.ProductID = product.ID
		if err := r.db.WithContext(ctx).Table("product_details").Save(&productDetail).Error; err != nil {
			// ROLLBACK ERROR WITH DELETE IMAGEKIT
			tx.Rollback()
			return entity.Product{}, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return entity.Product{}, err
	}

	var savedProduct entity.Product
	if err := r.db.WithContext(ctx).Preload("ProductImages").Preload("ProductDetails").First(&savedProduct, product.ID).Error; err != nil {
		return entity.Product{}, err
	}

	return savedProduct, nil
}

func (r *repository_product) GetProducts(ctx context.Context, limit int, offset int) ([]entity.Product, error) {
	products := []entity.Product{}

	query := r.db.WithContext(ctx).Preload("ProductImages").Preload("ProductDetails").Limit(limit).Offset(offset)

	if err := query.Find(&products).Error; err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}

func (r *repository_product) GetProductById(ctx context.Context, id uint64) (entity.Product, error) {
	product := entity.Product{}

	if err := r.db.WithContext(ctx).Preload("ProductImages").Preload("ProductDetails").Where("id = ?", id).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Product{}, errors.New("product not found")
		}

		return entity.Product{}, err
	}

	return product, nil
}

func (r *repository_product) ReadProductBySlug(ctx context.Context, slug string) (entity.Product, error) {
	var product entity.Product

	if err := r.db.WithContext(ctx).Preload("ProductImages").Preload("ProductDetails").Table("products").Where("slug = ?", slug).First(&product).Error; err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (r *repository_product) UpdateProduct(ctx context.Context, product *entity.Product) (entity.Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return entity.Product{}, err
	}

	return *product, nil
}

func (r *repository_product) DeleteProduct(ctx context.Context, product *entity.Product) (entity.Product, error) {
	if err := r.db.WithContext(ctx).Delete(&product).Error; err != nil {
		return entity.Product{}, err
	}

	return *product, nil
}
