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

func (h *galleryHandler) CreateGallery(c *gin.Context) {
	file, _ := c.FormFile("file")
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		fmt.Printf("error when open file %v", err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v", err)
		return
	}

	img, err := imagekits.Base64toEncode(buf.Bytes())
	if err != nil {
		fmt.Println("error reading image %v", err)
	}

	fmt.Println("image base 64 format : %v", img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var input gallery.InputGallery

	err = c.ShouldBind(&input)

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

	_, err = h.galleryService.CreateGallery(input, imageKitURL)
	if err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

func (h *galleryHandler) GetOneGallery(c *gin.Context) {
	var input gallery.InputGalleryID

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.galleryService.GetOneGallery(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDGallery /Gallery/GetByIDGallery", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, newDel)
	c.JSON(http.StatusOK, response)

}

func (h *galleryHandler) GetAllGallery(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newGalleryImage, err := h.galleryService.GetAllGallery(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "Eror")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDGallery /Gallery/GetByIDGallery", userAgent)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, newGalleryImage)
	c.JSON(http.StatusOK, response)
}

func (h *galleryHandler) UpdateGallery(c *gin.Context) {
	file, _ := c.FormFile("file")
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		fmt.Printf("error when open file %v", err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v", err)
		return
	}

	img, err := imagekits.Base64toEncode(buf.Bytes())
	if err != nil {
		fmt.Println("error reading image %v", err)
	}

	fmt.Println("image base 64 format : %v", img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var inputID gallery.InputGalleryID
	err = c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input gallery.InputGallery
	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	gallery, err := h.galleryService.UpdateGallery(inputID, input, imageKitURL)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, gallery)
	c.JSON(http.StatusOK, response)
}

func (h *galleryHandler) DeleteGallery(c *gin.Context) {
	var input gallery.InputGalleryID

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.galleryService.DeleteGallery(input.ID)
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
