package service

import (
	"context"
	"errors"
	"fmt"
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/formatter"
	"greenwelfare/repository"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ServiceProduct interface {
	CreateProduct(ctx context.Context, product dto.CreateProductDTO) (dto.ProductResponseDTO, error)
	GetProducts(ctx context.Context, limit int, offset int) ([]dto.ProductResponseDTO, error)
	ReadProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error)
	UpdateProductBySlug(ctx context.Context, slug string, product dto.UpdateProductDTO) (dto.ProductResponseDTO, error)
	DeleteProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error)
}

type service_product struct {
	repoProduct         repository.RepositoryProduct
	repoProductCategory repository.RepositoryProductCategory
}

func NewServiceProduct(repoProduct repository.RepositoryProduct, repoProductCategory repository.RepositoryProductCategory) *service_product {
	return &service_product{repoProduct: repoProduct, repoProductCategory: repoProductCategory}
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
			return dto.ProductResponseDTO{}, errors.New("invalid categoryId")
		}
		productCategory, err := s.repoProductCategory.GetProductCategoryById(ctx, uint64(categoryID))
		if err != nil {
			return dto.ProductResponseDTO{}, err
		}
		categoryID = int(productCategory.ID)
		isCategoryExist = true
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

	newProductImages := []entity.ProductImage{}
	for _, v := range product.ProductImages {
		productImage := entity.ProductImage{
			ImageUrl: v.ImageUrl,
		}
		newProductImages = append(newProductImages, productImage)
	}

	newProductDetails := []entity.ProductDetail{}
	for _, v := range product.ProductDetails {
		productDetail := entity.ProductDetail{
			Size:  v.Size,
			Price: v.Price,
			Stock: v.Stock,
		}
		newProductDetails = append(newProductDetails, productDetail)
	}

	res, errRepo := s.repoProduct.CreateProduct(ctx, newProduct, newProductImages, newProductDetails)
	if errRepo != nil {
		return dto.ProductResponseDTO{}, errRepo
	}
	return formatter.ParsingProductResponseDTO(res), nil
}

func (s *service_product) GetProducts(ctx context.Context, limit int, offset int) ([]dto.ProductResponseDTO, error) {
	res, errRepo := s.repoProduct.GetProducts(ctx, limit, offset)
	if errRepo != nil {
		return []dto.ProductResponseDTO{}, errRepo
	}

	responseDTO := []dto.ProductResponseDTO{}
	for _, v := range res {
		parsingRes := formatter.ParsingProductResponseDTO(v)
		responseDTO = append(responseDTO, parsingRes)
	}
	return responseDTO, nil
}

func (s *service_product) ReadProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error) {
	product, err := s.repoProduct.ReadProductBySlug(ctx, slug)

	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	return formatter.ParsingProductResponseDTO(product), nil
}

func (s *service_product) UpdateProductBySlug(
	ctx context.Context,
	slug string,
	newProduct dto.UpdateProductDTO,
) (dto.ProductResponseDTO, error) {
	product, err := s.repoProduct.ReadProductBySlug(ctx, slug)
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

	updated, err := s.repoProduct.UpdateProduct(ctx, &product)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	return formatter.ParsingProductResponseDTO(updated), nil
}

func (s *service_product) DeleteProductBySlug(ctx context.Context, slug string) (dto.ProductResponseDTO, error) {
	product, err := s.repoProduct.ReadProductBySlug(ctx, slug)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	newProduct, err := s.repoProduct.DeleteProduct(ctx, &product)
	if err != nil {
		return dto.ProductResponseDTO{}, err
	}

	return formatter.ParsingProductResponseDTO(newProduct), nil
}
