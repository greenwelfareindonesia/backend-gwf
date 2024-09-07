package handler

import (
	"bytes"
	"context"
	"fmt"
	"greenwelfare/dto"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"greenwelfare/service"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kajianHandler struct {
	kajianService service.ServiceKajian
}

func NewKajianHandler(kajianService service.ServiceKajian) *kajianHandler {
	return &kajianHandler{kajianService}
}

func (h *kajianHandler) CreateKajian(c *gin.Context) {
	var imagesKitURLs []string

	for i := 1; ; i++ {
		fileKey := fmt.Sprintf("file%d", i)
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

	var input dto.InputKajian
	if err := c.ShouldBind(&input); err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kajian, err := h.kajianService.Create(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	for _, url := range imagesKitURLs {
		if err := h.kajianService.CreateImage(kajian.ID, url); err != nil {
			response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
	}

	response := helper.SuccessfulResponse1(kajian)
	c.JSON(http.StatusCreated, response)
}

func (h *kajianHandler) GetOneKajian(c *gin.Context) {
	slug := c.Param("slug")

	kajian, err := h.kajianService.GetOne(slug)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(kajian)
	c.JSON(http.StatusOK, response)
}

func (h *kajianHandler) GetAllKajian(c *gin.Context) {
	kajians, err := h.kajianService.GetAll()
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(kajians)
	c.JSON(http.StatusOK, response)
}

func (h *kajianHandler) UpdateKajian(c *gin.Context) {
	slug := c.Param("slug")

	var update dto.UpdateKajian
	if err := c.ShouldBind(&update); err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kajian, err := h.kajianService.UpdateOne(slug, update)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(kajian)
	c.JSON(http.StatusOK, response)
}

func (h *kajianHandler) DeleteKajian(c * gin.Context) {
	slug := c.Param("slug")

	if err := h.kajianService.DeleteOne(slug); err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1("kajian successfully deleted")
	c.JSON(http.StatusOK, response)
}
