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

// @Summary Create New Kajian
// @Description Create New Kajian
// @Accept multipart/form-data
// @Security BearerAuth
// @Produce json
// @Tags Kajian
// @Param file1 formData file true "File gambar 1"
// @Param file2 formData file false "File gambar 2"
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/kajian [post]
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

	newKajian, err := h.kajianService.GetOneByID(kajian.ID);
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(newKajian)
	c.JSON(http.StatusCreated, response)
}

// @Summary Get One Kajian by slug
// @Description Get One Kajian by slug
// @Accept json
// @Produce json
// @Tags Kajian
// @Param slug path string true "Kajian slug"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/kajian/{slug} [get]
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

// @Summary Get All Kajian
// @Description Get All Kajian
// @Accept json
// @Produce json
// @Tags Kajian
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/kajian [get]
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

// @Summary Update Kajian
// @Description Update Kajian by slug
// @Accept multipart/form-data
// @Security BearerAuth
// @Produce json
// @Tags Kajian
// @Param file1 formData file false "File gambar 1"
// @Param file2 formData file false "File gambar 2"
// @Param slug path string true "Kajian Slug"
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/kajian/{slug} [put]
func (h *kajianHandler) UpdateKajian(c *gin.Context) {
	slug := c.Param("slug")

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

	var update dto.UpdateKajian
	if err := c.ShouldBind(&update); err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kajian, err := h.kajianService.UpdateOne(slug, update, imagesKitURLs)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(kajian)
	c.JSON(http.StatusOK, response)
}

// @Summary Delete Kajian by slug
// @Description Delete Kajian by slug
// @Accept json
// @Produce json
// @Tags Kajian
// @Security BearerAuth
// @Param slug path string true "Kajian slug"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/kajian/{slug} [delete]
func (h *kajianHandler) DeleteKajian(c *gin.Context) {
	slug := c.Param("slug")

	if err := h.kajianService.DeleteOne(slug); err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1("kajian successfully deleted")
	c.JSON(http.StatusOK, response)
}
