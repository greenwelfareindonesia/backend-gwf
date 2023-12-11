package handler

import (
	"greenwelfare/artikel"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type artikelHandler struct {
	artikelService artikel.Service
	endpointService endpointcount.StatisticsService
}

func NewArtikelHandler(artikelService artikel.Service, endpointService endpointcount.StatisticsService	) *artikelHandler {
	return &artikelHandler{artikelService, endpointService}
}

// @Summary Menghapus artikel by slug
// @Description Menghapus artikel by slug
// @Accept json
// @Produce json
// @Tags Artikel
// @Security BearerAuth
// @Param slug path int true "Artikel ID"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /artikel/{slug} [delete]
func (h *artikelHandler) DeleteArtikel(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.artikelService.DeleteArtikel(param)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.APIresponse(http.StatusOK, "article has successfully deleted")
	c.JSON(http.StatusOK, response)

}

// @Summary Mendapatkan satu artikel by slug
// @Description Mendapatkan satu artikel by slug
// @Accept json
// @Produce json
// @Tags Artikel
// @Param slug path int true "Artikel by slug"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /artikel/{slug} [get]
func (h *artikelHandler) GetOneArtikel(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.artikelService.GetOneArtikel(param)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
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

	response := helper.APIresponse(http.StatusOK, newDel)
	c.JSON(http.StatusOK, response)

}

// @Summary Memperbarui artikel
// @Description Memperbarui artikel dengan informasi yang diberikan
// @Accept multipart/form-data
// @Produce json
// @Tags Artikel
// @Security BearerAuth
// @Param slug path int true "Artikel slug"
// @Param File formData file true "File gambar"
// @Param FullName formData string true "Nama lengkap"
// @Param Email formData string true "Email"
// @Param Topic formData string true "Topik"
// @Param Message formData string true "Pesan artikel"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /artikel/{slug} [put]
func (h *artikelHandler) UpdateArtikel (c *gin.Context) {

	var input artikel.CreateArtikel

	param := c.Param("slug")

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	artikels, err := h.artikelService.UpdateArtikel(input, param)
	if err != nil {
		
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, artikel.UpdatedArticleFormat(artikels))
	c.JSON(http.StatusOK, response)
}

// @Summary Menambahkan artikel baru
// @Description Menambahkan entri artikel baru
// @Accept multipart/form-data
// @Produce json
// @Tags Artikel
// @Security BearerAuth
// @Param File formData file true "File gambar"
// @Param FullName formData string true "Nama lengkap"
// @Param Email formData string true "Email"
// @Param Topic formData string true "Topik"
// @Param Message formData string true "Pesan artikel"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /artikel [post]
func (h *artikelHandler) CreateArtikel(c *gin.Context) {
	var input artikel.CreateArtikel

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
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
	response := helper.APIresponse(http.StatusOK, artikel.PostArticleFormat(data))
	c.JSON(http.StatusOK, response)
}

// @Summary Mendapatkan semua artikel
// @Description Mendapatkan semua artikel
// @Accept json
// @Produce json
// @Tags Artikel
// @Param id query int false "ID"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /artikel [get]
func (h *artikelHandler) GetAllArtikel(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.artikelService.GetAllArtikel(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
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

	response := helper.APIresponse(http.StatusOK, newBerita)
	c.JSON(http.StatusOK, response)
}
