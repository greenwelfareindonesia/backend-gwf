package handler

import (
	"bytes"
	"context"
	"fmt"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"greenwelfare/workshop"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type workshopHandler struct {
	workshopService workshop.Service
	endpointService endpointcount.StatisticsService
}

func NewWorkshopHandler(workshopService workshop.Service, endpointService endpointcount.StatisticsService) *workshopHandler {
	return &workshopHandler{workshopService, endpointService}
}

// @Summary Create New Workshop
// @Description Create New Workshop
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Workshop
// @Param File formData file true "File"
// @Param Title formData string true "Title"
// @Param Description formData string true "Description"
// @Param Date formData string true "Date"
// @Param Url formData string true "Url"
// @Param IsOpen formData boolean true "IsOpen"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/workshop [post]
func (h *workshopHandler) CreateWorkshop(c *gin.Context) {
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

	var input workshop.CreateWorkshop

	err = c.ShouldBind(&input)

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

	data, err := h.workshopService.CreateWorkshop(input, imageKitURL)
	if err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	// data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, workshop.PostFormatterWorkshop(data))
	c.JSON(http.StatusOK, response)
}

// @Summary Get One Workshop by slug 
// @Description Get One Workshop by slug 
// @Accept json
// @Produce json
// @Tags Workshop
// @Param slug path string true "slug Workshop"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/workshop/{slug} [get]
func (h *workshopHandler) GetOneWorkshop(c *gin.Context) {
	param := c.Param("slug")

	err := c.ShouldBindUri(&param)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.workshopService.GetOneWorkshop(param)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDWorkshop /Workshop/GetByIDWorkshop", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, (newDel))
	c.JSON(http.StatusOK, response)

}

// @Summary Get All Workshop
// @Description Get All Workshop
// @Accept json
// @Produce json
// @Tags Workshop
// @Param id query int false "ID Workshop"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/workshop [get]
func (h *workshopHandler) GetAllWorkshop(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.workshopService.GetAllWorkshop(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetAllWorkshop /Workshop/GetAllWorkshop", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, newBerita)
	c.JSON(http.StatusOK, response)
}

// @Summary Update workshop by slug
// @Description Update workshop by slug 
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Workshop
// @Param slug path string true "slug Workshop"
// @Param File formData file true "File"
// @Param Title formData string true "Title"
// @Param Description formData string true "Description"
// @Param Date formData string true "Date"
// @Param Url formData string true "Url"
// @Param IsOpen formData boolean true "IsOpen"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/workshop/{slug} [put]
func (h *workshopHandler) UpdateWorkshop(c *gin.Context) {
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
	err = c.ShouldBindUri(&param)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input workshop.CreateWorkshop
	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.workshopService.UpdateWorkshop(param, input, imageKitURL)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// data := gin.H{"is_updated": true}
	response := helper.APIresponse(http.StatusOK, workshop.UpdateFormatterWorkshop(data))
	c.JSON(http.StatusOK, response)
}

// @Summary Delete workshop by slug
// @Description Delete workshop by slug
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Workshop
// @Param slug path string true "slug Workshop"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/workshop/{slug} [delete]
func (h *workshopHandler) DeleteWorkshop(c *gin.Context) {
	param := c.Param("slug")

	err := c.ShouldBindUri(&param)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.workshopService.DeleteWorkshop(param)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.APIresponse(http.StatusOK, "workshop has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
