package handler

import (
	"greenwelfare/dto"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/helper"
	"greenwelfare/service"
	"net/http"

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
