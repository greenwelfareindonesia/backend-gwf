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

// @Summary Menghapus artikel by id
// @Description Menghapus artikel by id
// @Accept json
// @Produce json
// @Tags Artikel
// @Security BearerAuth
// @Param id path int true "Artikel ID"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /artikel/{id} [delete]
func (h *artikelHandler) DeleteArtikel(c *gin.Context) {
	var input artikel.GetArtikel

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.artikelService.DeleteArtikel(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.APIresponse(http.StatusOK, newDel)
	c.JSON(http.StatusOK, response)

}

// @Summary Mendapatkan satu artikel by id
// @Description Mendapatkan satu artikel by id
// @Accept json
// @Produce json
// @Tags Artikel
// @Param id path int true "Artikel ID"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /artikel/{id} [get]
func (h *artikelHandler) GetOneArtikel(c *gin.Context) {
	var input artikel.GetArtikel

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.artikelService.GetOneArtikel(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDArtikel /Artikel/GetByIDArtikel", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
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
// @Param id path int true "Artikel ID"
// @Param file formData file true "File gambar"
// @Param full_name formData string true "Nama lengkap"
// @Param email formData string true "Email"
// @Param topic formData string true "Topik"
// @Param message formData string true "Pesan artikel"
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{}
// @Success 422 {object} map[string]interface{}
// @Success 500 {object} map[string]interface{}
// @Router /artikel/{id} [put]
func (h *artikelHandler) UpdateArtikel (c *gin.Context) {
	var inputID artikel.GetArtikel

	var input artikel.CreateArtikel

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	artikel, err := h.artikelService.UpdateArtikel(input, inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, artikel)
	c.JSON(http.StatusOK, response)
}

// @Summary Menambahkan artikel baru
// @Description Menambahkan entri artikel baru
// @Accept multipart/form-data
// @Produce json
// @Tags Artikel
// @Security BearerAuth
// @Param file formData file true "File gambar"
// @Param full_name formData string true "Nama lengkap"
// @Param email formData string true "Email"
// @Param topic formData string true "Topik"
// @Param message formData string true "Pesan artikel"
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
		//inisiasi data yang tujuan dalam return hasil ke postman
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.artikelService.CreateArtikel(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
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
		response := helper.APIresponse(http.StatusUnprocessableEntity, "Eror")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetAllArtikel /Artikel/GetAllArtikel", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusOK, newBerita)
	c.JSON(http.StatusOK, response)
}
