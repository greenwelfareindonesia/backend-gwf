package formatter

import (
	"greenwelfare/dto"
	"greenwelfare/entity"
)

func ParsingProductResponseDTO(product entity.Product) dto.ProductResponseDTO {
	response := dto.ProductResponseDTO{
		ID:          product.ID,
		Slug:        product.Slug,
		Name:        product.Name,
		Excerpt:     product.Excerpt,
		Description: product.Description,
		Merk:        product.Merk,
		ItemWeight:  product.ItemWeight,
		IsActive:    product.IsActive,
	}

	if len(product.ProductImages) > 0 {
		for _, v := range product.ProductImages {
			productImage := dto.ProductImageResponseDTO{
				ID:        v.ID,
				ProductID: v.ProductID,
				ImageUrl:  v.ImageUrl,
			}
			response.ProductImages = append(response.ProductImages, &productImage)
		}
	}

	if len(product.ProductDetails) > 0 {
		for _, v := range product.ProductDetails {
			productDetail := dto.ProductDetailResponseDTO{
				ID:    v.ID,
				Size:  v.Size,
				Price: v.Price,
				Stock: v.Stock,
			}
			response.ProductDetails = append(response.ProductDetails, &productDetail)
		}
	}

	return response
}