package handler

import (
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/dto"
	"greenwelfare/formatter"
	"greenwelfare/service"
	"greenwelfare/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type artikelHandler struct {
	artikelService  service.ServiceArtikel
	endpointService endpointcount.StatisticsService
}

func NewArtikelHandler(artikelService service.ServiceArtikel, endpointService endpointcount.StatisticsService) *artikelHandler {
	return &artikelHandler{artikelService, endpointService}
}

// @Summary Delete article by slug
// @Description Delete article by slug
// @Accept json
// @Produce json
// @Tags Article
// @Security BearerAuth
// @Param slug path string true "Article By Slug"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /api/article/{slug} [delete]
func (h *artikelHandler) DeleteArtikel(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.artikelService.DeleteArtikel(param)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.SuccessfulResponse1("article has successfully deleted")
	c.JSON(http.StatusOK, response)

}

// @Summary Get one article by slug
// @Description get one article by slug
// @Accept json
// @Produce json
// @Tags Article
// @Security BearerAuth
// @Param slug path string true "Article by slug"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /api/article/{slug} [get]
func (h *artikelHandler) GetOneArtikel(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.artikelService.GetOneArtikel(param)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDArtikel /Artikel/GetByIDArtikel", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(newDel)
	c.JSON(http.StatusOK, response)

}

// @Summary update article
// @Description update article
// @Accept multipart/form-data
// @Produce json
// @Tags Article
// @Security BearerAuth
// @Param slug path string true "Article by slug"
// @Param fullName formData string true "Full Name"
// @Param email formData string true "Email"
// @Param topic formData string true "Topic"
// @Param articleMessage formData string true "Article Message"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /api/article/{slug} [put]
func (h *artikelHandler) UpdateArtikel(c *gin.Context) {

	var input dto.CreateArtikel

	param := c.Param("slug")

	err := c.ShouldBind(&input)
	if err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	artikels, err := h.artikelService.UpdateArtikel(input, param)
	if err != nil {

		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), artikels)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(formatter.UpdatedArticleFormat(artikels))
	c.JSON(http.StatusOK, response)
}

// @Summary Create new article
// @Description Create new article
// @Accept multipart/form-data
// @Produce json
// @Tags Article
// @Security BearerAuth
// @Param fullName formData string true "Full Name"
// @Param email formData string true "Email"
// @Param topic formData string true "Topic"
// @Param articleMessage formData string true "Article Message"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /api/article [post]
func (h *artikelHandler) CreateArtikel(c *gin.Context) {
	var input dto.CreateArtikel

	err := c.ShouldBind(&input)

	if err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.artikelService.CreateArtikel(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// data := gin.H{"is_uploaded": true}
	response := helper.SuccessfulResponse1(formatter.PostArticleFormat(data))
	c.JSON(http.StatusOK, response)
}

// @Summary get all article
// @Description get all article
// @Accept json
// @Produce json
// @Tags Article
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /api/article [get]
func (h *artikelHandler) GetAllArtikel(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.artikelService.GetAllArtikel(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetAllArtikel /Artikel/GetAllArtikel", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(newBerita)
	c.JSON(http.StatusOK, response)
}
