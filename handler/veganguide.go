package handler

import (
	"bytes"
	"context"
	"fmt"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"greenwelfare/veganguide"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type veganguideHandler struct {
	veganguideService veganguide.Service
	endpointService   endpointcount.StatisticsService
}

func NewVeganguideHandler(veganguideService veganguide.Service, endpointService endpointcount.StatisticsService) *veganguideHandler {
	return &veganguideHandler{veganguideService, endpointService}
}

// @Summary Hapus data Veganguide berdasarkan ID
// @Description Hapus data Veganguide berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags VeganGuide
// @Param id path int true "ID Veganguide"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /veganguide/{id} [delete]
func (h *veganguideHandler) DeleteVeganguide(c *gin.Context) {
	var input veganguide.GetVeganguide

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.veganguideService.DeleteVeganguide(input.ID)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)
}

// @Summary Dapatkan Veganguide berdasarkan ID
// @Description Dapatkan Veganguide berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Tags VeganGuide
// @Param id path int true "ID Veganguide"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /veganguide/{id} [get]
func (h *veganguideHandler) GetVeganguideByID(c *gin.Context) {
	// var input veganguide.GetVeganguide
	param := c.Param("slug")

	// err := c.ShouldBindUri(&input)
	data, err := h.veganguideService.GetOneVeganguide(param)

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

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDVeganguide /Veganguide/GetByIDVeganguide", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)

}

// @Summary Dapatkan semua Veganguide atau Veganguide berdasarkan ID tertentu
// @Description Dapatkan semua Veganguide atau Veganguide berdasarkan ID tertentu
// @Accept json
// @Produce json
// @Tags VeganGuide
// @Param id query int false "ID Veganguide"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /veganguide [get]
func (h *veganguideHandler) GetAllVeganguide(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	data, err := h.veganguideService.GetAllVeganguide(input)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetAllVeganguide /Veganguide/GetAllVeganguide", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)
}

// @Summary Buat data Veganguide baru
// @Description Buat data Veganguide baru dengan informasi yang diberikan
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags VeganGuide
// @Param file formData file true "File gambar"
// @Param judul formData string true "Judul"
// @Param deskripsi formData string true "Deskripsi"
// @Param body formData string false "Body"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /veganguide [post]
func (h *veganguideHandler) PostVeganguideHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	src, err := file.Open()
	if err != nil {
		fmt.Printf("error when open file %v", err)
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v", err)
		return
	}

	img, err := imagekits.Base64toEncode(buf.Bytes())
	if err != nil {
		fmt.Println("error reading image ", err)
	}

	fmt.Println("image base 64 format : ", img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var input veganguide.VeganguideInput

	err = c.ShouldBind(&input)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errorMessages}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		//inisiasi data yang tujuan dalam return hasil ke postman
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.veganguideService.CreateVeganguide(input, imageKitURL)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

// @Summary Update data Veganguide berdasarkan ID
// @Description Update data Veganguide berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags VeganGuide
// @Param id path int true "ID Veganguide"
// @Param file formData file true "File gambar"
// @Param judul formData string true "Judul"
// @Param deskripsi formData string true "Deskripsi"
// @Param body formData string false "Body"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /veganguide/{id} [put]
func (h *veganguideHandler) UpdateVeganguide(c *gin.Context) {
	file, _ := c.FormFile("file")
	src, err := file.Open()
	if err != nil {
		fmt.Printf("error when open file %v", err)
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v", err)
		return
	}

	img, err := imagekits.Base64toEncode(buf.Bytes())
	if err != nil {
		fmt.Println("error reading image ", err)
	}

	fmt.Println("image base 64 format : ", img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	param := c.Param("slug")

	var input veganguide.VeganguideInput
	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	veganguide, err := h.veganguideService.UpdateVeganguide(input, param, imageKitURL)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, veganguide)
	c.JSON(http.StatusOK, response)
}
