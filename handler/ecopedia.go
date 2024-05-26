package handler

import (
	"bytes"
	"context"
	"fmt"
	"greenwelfare/dto"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/formatter"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"greenwelfare/service"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ecopediaHandler struct {
	ecopediaService service.ServiceEcopedia
	endpointService endpointcount.StatisticsService
}

func NewEcopediaHandler(ecopediaService service.ServiceEcopedia, endpointService endpointcount.StatisticsService) *ecopediaHandler {
	return &ecopediaHandler{ecopediaService, endpointService}
}

// @Summary Delete Ecopedia by id
// @Description Delete Ecopedia by id
// @Accept json
// @Produce json
// @Tags Ecopedia
// @Security BearerAuth
// @Param id path int true "Ecopedia id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ecopedia/{id} [delete]
func (h *ecopediaHandler) DeleteEcopedia(c *gin.Context) {
	var inputID dto.EcopediaID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), inputID)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.ecopediaService.DeleteEcopedia(inputID.ID)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), inputID)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1("ecopedia has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}

// @Summary Get One Ecopedia by slug
// @Description Get One Ecopedia by slug
// @Accept json
// @Produce json
// @Tags Ecopedia
// @Param slug path string true "Ecopedia slug"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ecopedia/{slug} [get]
func (h *ecopediaHandler) GetEcopediaByID(c *gin.Context) {
	param := c.Param("slug")

	data, err := h.ecopediaService.GetEcopediaByID(param)
	if err != nil {

		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDEcopedia /Ecopedia/GetByIDEcopedia", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(formatter.GetOneEcopediaFormat(data))
	c.JSON(http.StatusOK, response)

}

// @Summary Get All Ecopedia
// @Description Get All Ecopedia
// @Accept json
// @Produce json
// @Tags Ecopedia
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ecopedia [get]
func (h *ecopediaHandler) GetAllEcopedia(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	data, err := h.ecopediaService.GetAllEcopedia(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetAllEcopedia /Ecopedia/GetAllEcopedia", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(formatter.FormatterGetAllEcopedia(data))
	c.JSON(http.StatusOK, response)
}

// @Summary Update Ecopedia
// @Description Update Ecopedia by slug
// @Accept multipart/form-data
// @Security BearerAuth
// @Produce json
// @Tags Ecopedia
// @Param slug path string true "Ecopedia Slug"
// @Param Title formData string true "Title"
// @Param SubTitle formData string true "SubTitle"
// @Param Description formData string true "Description"
// @Param SrcFile formData string true "SrcFile"
// @Param Reference formData string true "Reference"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ecopedia/{slug} [put]
func (h *ecopediaHandler) UpdateEcopedia(c *gin.Context) {
	var input dto.EcopediaInput

	param := c.Param("slug")

	err := c.ShouldBind(&input)
	if err != nil {
		// errorMessages := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errorMessages}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	_, err = h.ecopediaService.UpdateEcopedia(param, input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_updated": true}
	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)

}

// @Summary Create New Ecopedia
// @Description Create New Ecopedia
// @Accept multipart/form-data
// @Security BearerAuth
// @Produce json
// @Tags Ecopedia
// @Param File1 formData file true "File gambar 1"
// @Param File2 formData file true "File gambar 2"
// @Param Title formData string true "Title"
// @Param SubTitle formData string true "SubTitle"
// @Param Description formData string true "Description"
// @Param SrcFile formData string true "SrcFile"
// @Param Reference formData string true "Reference"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/ecopedia [post]
func (h *ecopediaHandler) PostEcopediaHandler(c *gin.Context) {
	var imagesKitURLs []string

	for i := 1; ; i++ {
		fileKey := fmt.Sprintf("File%d", i)
		file, err := c.FormFile(fileKey)

		// If there are no more files to upload, break the loop
		if err == http.ErrMissingFile {
			break
		}

		if err != nil {
			fmt.Printf("Error when opening file %s: %v\n", fileKey, err)
			continue // Skip to the next file
		}

		src, err := file.Open()
		if err != nil {
			fmt.Printf("Error when opening file %s: %v\n", fileKey, err)
			continue
		}
		defer src.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, src); err != nil {
			fmt.Printf("Error reading file %s: %v\n", fileKey, err)
			continue
		}

		img, err := imagekits.Base64toEncode(buf.Bytes())
		if err != nil {
			fmt.Printf("Error reading image %s: %v\n", fileKey, err)
			continue
		}

		fmt.Printf("Image base64 format %s: %v\n", fileKey, img)

		imageKitURL, err := imagekits.ImageKit(context.Background(), img)
		if err != nil {
			fmt.Printf("Error uploading image %s to ImageKit: %v\n", fileKey, err)
			continue
		}

		imagesKitURLs = append(imagesKitURLs, imageKitURL)
	}

	var ecopediaInput dto.EcopediaInput

	err := c.ShouldBind(&ecopediaInput)

	if err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), ecopediaInput)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Create a new news item with the provided input
	newNews, err := h.ecopediaService.CreateEcopedia(ecopediaInput)
	// fmt.Println(newNews)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), newNews)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Associate the uploaded images with the news item
	for _, imageURL := range imagesKitURLs {
		// Create a new BeritaImage record for each image and associate it with the news item
		err := h.ecopediaService.CreateEcopediaImage(newNews.ID, imageURL)
		if err != nil {
			response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	data := gin.H{"is_uploaded": true}
	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)

}
