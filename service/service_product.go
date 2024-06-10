package service

import (
	"context"
	"fmt"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/repository"
	"math/rand"
	"strings"
	"time"
)

type ServiceProduct interface {
	CreateProduct(ctx context.Context, product dto.CreateProductDTO) (dto.ProductResponseDTO, error)
	ReadProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error)
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
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	randomNumber := seededRand.Intn(1000000)

	newProduct := entity.Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
		Slug:        fmt.Sprintf("%s-%d", slugName, randomNumber),
	}

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

func parsingProductResponseDTO(product entity.Product) dto.ProductResponseDTO {
	response := dto.ProductResponseDTO{
		ID:            product.ID,
		Name:          product.Name,
		Slug:          product.Slug,
		Stock:         product.Stock,
		Description:   product.Description,
		DefaultColumn: product.DefaultColumn,
	}
	return response
}
