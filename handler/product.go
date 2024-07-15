package handler

import (
	"bytes"
	"context"
	// "context"
	"fmt"
	"greenwelfare/dto"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"greenwelfare/service"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService  service.ServiceProduct
	endpointService endpointcount.StatisticsService
}

func NewProductHandler(svc service.ServiceProduct, endpointService endpointcount.StatisticsService) *productHandler {
	return &productHandler{svc, endpointService}
}

// @Summary Create new product
// @Description Create new product
// @Accept multipart/form-data
// @Produce json
// @Tags Product
// @Param Name formData string true "Name"
// @Param Price formData string true "Price"
// @Param Description formData string true "Description"
// @Param ImageUrl formData file true "File gambar product"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /api/product  [post]
func (h *productHandler) CreateProduct(ctx *gin.Context) {
	newProductRequest := dto.CreateProductDTO{}
	var errBinding []string

	// BINDING FORM DATA
	if err := ctx.Bind(&newProductRequest); err != nil {
		errBindJson := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), newProductRequest)
		ctx.AbortWithStatusJSON(errBindJson.Error.Code, errBindJson)
		return
	}

	// VALIDATION INPUT VALUE
	if err := newProductRequest.Validate(); err != nil {
		errValidate := helper.FailedResponse1(http.StatusBadRequest, err.Error(), newProductRequest)
		ctx.AbortWithStatusJSON(errValidate.Error.Code, errValidate)
		return
	}

	// BINDING PRODUCT DETAIL
	var productDetails []dto.CreateProductDetailDTO
	for i := 0; ; i++ {
		var newProductDetail dto.CreateProductDetailDTO

		size := ctx.PostForm(fmt.Sprintf("product_details[%d][size]", i))
		price := ctx.PostForm(fmt.Sprintf("product_details[%d][price]", i))
		stock := ctx.PostForm(fmt.Sprintf("product_details[%d][stock]", i))

		if size == "" && price == "" && stock == "" {
			break
		}

		priceToInt, errPrice := strconv.Atoi(price)
		if errPrice != nil {
			errBinding = append(errBinding, fmt.Sprintf("invalid input price on product_details[%d][size]", i))
			continue
		}
		stockToInt, errStock := strconv.Atoi(stock)
		if errStock != nil {
			errBinding = append(errBinding, fmt.Sprintf("invalid stock price on product_details[%d][stock]", i))
			continue
		}

		newProductDetail.Size = size
		newProductDetail.Price = uint64(priceToInt)
		newProductDetail.Stock = uint64(stockToInt)

		// update total stock
		newProductRequest.TotalStock += uint64(stockToInt)

		if err := newProductDetail.Validate(); err != nil {
			errBinding = append(errBinding, fmt.Sprintf("invalid input product_detail[%d] : %s", i, err.Error()))
			continue
		}

		productDetails = append(productDetails, newProductDetail)
	}
	newProductRequest.ProductDetails = productDetails

	// CHECK ERROR BINDING PRODUCT DETAIL
	if len(errBinding) > 0 {
		errValidate := helper.FailedResponse2(http.StatusBadRequest, errBinding)
		ctx.AbortWithStatusJSON(400, errValidate)
		return
	}

	// CHECK PRODUCT DETAIL DATA
	if len(productDetails) < 1 {
		errProductDetails := helper.FailedResponse1(http.StatusUnprocessableEntity, "product_details is required", newProductRequest)
		ctx.AbortWithStatusJSON(errProductDetails.Error.Code, errProductDetails)
		return
	}

	// PROCESS IMAGE UPLOAD
	var productImages []dto.CreateProductImageDTO
	for i := 0; ; i++ {
		var productImg dto.CreateProductImageDTO
		fileKey := fmt.Sprintf("product_images[%d]", i)
		imageHeader, err := ctx.FormFile(fileKey)
		if err != nil {
			break
		}

		// VALIDATE MIMETYPES
		mimeType := imageHeader.Header.Get("Content-Type")
		if mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/jpg" {
			var errVal = fmt.Sprintf("%s File format not allowed", fileKey)
			errBinding = append(errBinding, errVal)
			fmt.Println(errVal)
			continue
		}

		// OPEN THE FILE
		file, err := imageHeader.Open()
		if err != nil {
			var errVal = fmt.Sprintf("Error when opening file with key [%s]: %v\n", fileKey, err.Error())
			errBinding = append(errBinding, errVal)
			fmt.Println(errVal)
			continue
		}

		defer file.Close()

		// UPLOAD IMAGE
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			var errVal = fmt.Sprintf("Error buffering image %s: %v\n", fileKey, err.Error())
			errBinding = append(errBinding, errVal)
			fmt.Println(errVal)
			continue
		}

		img, err := imagekits.Base64toEncode(buf.Bytes())
		if err != nil {
			var errVal = fmt.Sprintf("Error encoding image to base64 with key [%s]: %v\n", fileKey, err.Error())
			errBinding = append(errBinding, errVal)
			fmt.Println(errVal)
			continue
		}

		imageKitUrl, err := imagekits.ImageKit(context.Background(), img)
		if err != nil {
			var errVal = fmt.Sprintf("Error upload image with key [%s]: %v\n", fileKey, err.Error())
			errBinding = append(errBinding, errVal)
			fmt.Println(errVal)
			continue
		}
		productImg.ImageUrl = imageKitUrl
		productImages = append(productImages, productImg)
	}

	// CHECK ERROR BINDING PRODUCT IMAGE
	if len(errBinding) > 0 {
		// DELETE IMAGE FROM IMAGEKIT

		errValidate := helper.FailedResponse2(http.StatusBadRequest, errBinding)
		ctx.AbortWithStatusJSON(400, errValidate)
		return
	}

	// CHECK PRODUCT IMAGES DATA
	if len(productImages) < 1 {
		errProductImages := helper.FailedResponse1(http.StatusUnprocessableEntity, "product_images is required", newProductRequest)
		ctx.AbortWithStatusJSON(errProductImages.Error.Code, errProductImages)
		return
	}

	newProductRequest.ProductImages = productImages
	product, errSvc := h.productService.CreateProduct(ctx, newProductRequest)
	if errSvc != nil {
		errCreate := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), newProductRequest)
		ctx.AbortWithStatusJSON(errCreate.Error.Code, errCreate)
		return
	}

	response := helper.SuccessfulResponse1(product)
	ctx.JSON(http.StatusOK, response)

}

func (h *productHandler) ReadProductBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	product, err := h.productService.ReadProductBySlug(ctx, slug)

	if err != nil {
		err := helper.FailedResponse1(http.StatusNotFound, err.Error(), nil)
		ctx.AbortWithStatusJSON(err.Error.Code, err)
		return
	}

	response := helper.SuccessfulResponse1(product)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProductBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	var input dto.UpdateProductDTO
	if err := ctx.ShouldBind(&input); err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	product, err := h.productService.UpdateProductBySlug(ctx, slug, input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), product)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(product)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProductBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if _, err := h.productService.DeleteProductBySlug(ctx, slug); err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), slug)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1("product has successfully deleted")
	ctx.JSON(http.StatusOK, response)
}
