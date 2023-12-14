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

// @Summary Delete Veganguide by slug
// @Description Delete Veganguide by slug
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags VeganGuide
// @Param slug path string true "slug Veganguide"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/veganguide/{slug} [delete]
func (h *veganguideHandler) DeleteVeganguide(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.veganguideService.DeleteVeganguide(param)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, "veganguide has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}

// @Summary Get One Veganguide by slug
// @Description Get One Veganguide by slug
// @Accept json
// @Produce json
// @Tags VeganGuide
// @Param slug path string true "slug Veganguide"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/veganguide/{slug} [get]
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

// @Summary Get All Veganguide
// @Description Get All Veganguide
// @Accept json
// @Produce json
// @Tags VeganGuide
// @Param slug query string false "slug Veganguide"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/veganguide [get]
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

// @Summary Create New Veganguide 
// @Description Create New Veganguide
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags VeganGuide
// @Param File formData file true "File"
// @Param Judul formData string true "Judul"
// @Param Deskripsi formData string true "Deskripsi"
// @Param body formData string false "Body"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/veganguide [post]
func (h *veganguideHandler) PostVeganguideHandler(c *gin.Context) {
	file, _ := c.FormFile("File")
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

	data, err := h.veganguideService.CreateVeganguide(input, imageKitURL)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	
	response := helper.APIresponse(http.StatusOK, veganguide.PostFormatterWorkshop(data))
	c.JSON(http.StatusOK, response)
}

// @Summary Update Veganguide by slug
// @Description Update Veganguide by slug
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags VeganGuide
// @Param slug path string true "slug Veganguide"
// @Param File formData file true "File gambar"
// @Param Judul formData string true "Judul"
// @Param Deskripsi formData string true "Deskripsi"
// @Param Body formData string false "Body"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/veganguide/{slug} [put]
func (h *veganguideHandler) UpdateVeganguide(c *gin.Context) {
	file, _ := c.FormFile("File")
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

	veganguides, err := h.veganguideService.UpdateVeganguide(input, param, imageKitURL)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, veganguide.UpdatedFormatterWorkshop(veganguides))
	c.JSON(http.StatusOK, response)
}
