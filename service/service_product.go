package service

import (
	"context"
	"errors"
	"fmt"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ServiceProduct interface {
	CreateProduct(ctx context.Context, product dto.CreateProductDTO) (dto.ProductResponseDTO, error)
	ReadProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error)
	UpdateProductBySlug(ctx context.Context, slug string, product dto.UpdateProductDTO) (dto.ProductResponseDTO, error)
	DeleteProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error)
}

type service_product struct {
	repository repository.RepositoryProduct
}

func NewServiceProduct(repository repository.RepositoryProduct) *service_product {
	return &service_product{repository: repository}
}

func (s *service_product) CreateProduct(ctx context.Context, product dto.CreateProductDTO) (dto.ProductResponseDTO, error) {
	slugName := strings.ToLower(product.Name)
	slugName = strings.ReplaceAll(slugName, " ", "-")
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := seededRand.Intn(1000000)

	itemWeight, err := strconv.ParseFloat(product.ItemWeight, 64)
	if err != nil {
		return dto.ProductResponseDTO{}, errors.New("invalid item_weight")
	}

	var categoryID = 0
	var isCategoryExist = false
	if product.CategoryID != "" {
		categoryID, err := strconv.Atoi(product.CategoryID)
		if err != nil {
			return dto.ProductResponseDTO{}, errors.New("invalid category_id")
		}
		_ = categoryID
		// find category
	}

	newProduct := entity.Product{
		Name:        product.Name,
		Slug:        fmt.Sprintf("%s-%d", slugName, randomNumber),
		Excerpt:     product.Excerpt,
		Description: product.Description,
		Merk:        product.Merk,
		TotalStock:  product.TotalStock,
		TotalSales:  0,
		ItemWeight:  itemWeight,
	}

	if categoryID > 0 && isCategoryExist {
		newProduct.CategoryID = uint64(categoryID)
	}

	// USING DB TRANSACTIONAL
	// CREATE PRODUCT


	// CREATE PRODUCT IMAGE


	// CREATE PRODUCT DETAIL

	// ROLLBACK ERROR WITH DELETE IMAGEKIT

	res, errRepo := s.repository.CreateProduct(ctx, newProduct)
	if errRepo != nil {
		return dto.ProductResponseDTO{}, errRepo
	}
	return parsingProductResponseDTO(res), nil
}

func (s *service_product) ReadProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error) {
	product, err := s.repository.ReadProductBySlug(ctx, slug)

	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	return parsingProductResponseDTO(product), nil
}

func (s *service_product) UpdateProductBySlug(
	ctx context.Context,
	slug string,
	newProduct dto.UpdateProductDTO,
) (dto.ProductResponseDTO, error) {
	product, err := s.repository.ReadProductBySlug(ctx, slug)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	slugName := strings.ToLower(product.Name)
	slugName = strings.ReplaceAll(slugName, " ", "-")
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := seededRand.Intn(1000000)

	product.Slug = fmt.Sprintf("%s-%d", slugName, randomNumber)
	product.Name = newProduct.Name
	// product.Price = newProduct.Price
	// product.Description = newProduct.Description
	// product.Stock = newProduct.Stock

	updated, err := s.repository.UpdateProduct(ctx, &product)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	return parsingProductResponseDTO(updated), nil
}

func (s *service_product) DeleteProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error) {
	product, err := s.repository.ReadProductBySlug(ctx, slug)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	newProduct, err := s.repository.DeleteProduct(ctx, &product)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	return parsingProductResponseDTO(newProduct), nil
}

func parsingProductResponseDTO(product entity.Product) dto.ProductResponseDTO {
	response := dto.ProductResponseDTO{
		ID:   product.ID,
		Name: product.Name,
		Slug: product.Slug,
		// Stock:         product.Stock,
		Description:   product.Description,
		DefaultColumn: product.DefaultColumn,
	}
	return response
}
