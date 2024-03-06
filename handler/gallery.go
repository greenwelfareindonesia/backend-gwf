package handler

import (
	"bytes"
	"context"
	"fmt"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/gallery"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type galleryHandler struct {
	galleryService  gallery.Service
	endpointService endpointcount.StatisticsService
}

func NewGalleryHandler(galleryService gallery.Service, endpointService endpointcount.StatisticsService) *galleryHandler {
	return &galleryHandler{galleryService, endpointService}
}

// @Summary Create New Gallery 
// @Description Create New Gallery 
// @Accept json
// @Produce json
// @Tags Gallery
// @Security BearerAuth
// @Param File1 formData file true "File gambar"
// @Param File2 formData file true "File gambar"
// @Param Alt formData string true "Alt"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/gallery [post]
func (h *galleryHandler) CreateGallery(c *gin.Context) {
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

		var ecopediaInput gallery.InputGallery

		err := c.ShouldBind(&ecopediaInput)

		if err != nil {
			// errors := helper.FormatValidationError(err)
			// errorMessage := gin.H{"errors": errors}
			response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), ecopediaInput)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		// Create a new news item with the provided input
		newNews, err := h.galleryService.CreateGallery(ecopediaInput)
		fmt.Println(newNews)
		if err != nil {
			response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), newNews)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		// Associate the uploaded images with the news item
		for _, imageURL := range imagesKitURLs {
			// Create a new BeritaImage record for each image and associate it with the news item
			err := h.galleryService.CreateImageGallery(newNews.ID, imageURL)
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

// @Summary Get One Gallery by slug
// @Description Get One Gallery by slug 
// @Accept json
// @Produce json
// @Tags Gallery
// @Param slug path string true "slug Gallery"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/gallery/{slug} [get]
func (h *galleryHandler) GetOneGallery(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.galleryService.GetOneGallery(param)
	if err != nil {
		
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDGallery /Gallery/GetByIDGallery", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(gallery.PostFormatterGallery(newDel))
	c.JSON(http.StatusOK, response)

}

// @Summary Get All Gallery 
// @Description Get All Gallery 
// @Accept json
// @Produce json
// @Tags Gallery
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/gallery [get]
func (h *galleryHandler) GetAllGallery(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newGalleryImage, err := h.galleryService.GetAllGallery(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDGallery /Gallery/GetByIDGallery", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(gallery.FormatterGetAllGallery(newGalleryImage))
	c.JSON(http.StatusOK, response)
}

// @Summary Update Gallery by Slug
// @Description Update Gallery by Slug 
// @Accept json
// @Produce json
// @Tags Gallery
// @Security BearerAuth
// @Param slug path int true "Slug Gallery"
// @Param Alt formData string true "Alt"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/gallery/{slug} [put]
func (h *galleryHandler) UpdateGallery(c *gin.Context) {

	param := c.Param("slug")

	var input gallery.InputGallery
	err := c.ShouldBind(&input)
	if err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.galleryService.UpdateGallery(param, input)
	if err != nil {
		
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_updated": true}
	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)
}

// @Summary Delete Gallery by id
// @Description Delete Gallery by id 
// @Accept json
// @Produce json
// @Tags Gallery
// @Security BearerAuth
// @Param id path int true "id Gallery"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/gallery/{id} [delete]
func (h *galleryHandler) DeleteGallery(c *gin.Context) {
	param := c.Param("id")
	conv, _ := strconv.Atoi(param)

	_, err := h.galleryService.DeleteGallery(conv)
	if err != nil {
		
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.SuccessfulResponse1("gallery has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}
