package handler

import (
	"bytes"
	"context"
	"fmt"
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
}

func NewWorkshopHandler(workshopService workshop.Service) *workshopHandler {
	return &workshopHandler{workshopService}
}

func (h *workshopHandler) CreateWorkshop(c *gin.Context) {
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
		//inisiasi data yang tujuan dalam return hasil ke postman
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.workshopService.CreateWorkshop(input, imageKitURL)
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

func (h *workshopHandler) GetOneWorkshop(c *gin.Context) {
	var input workshop.GetWorkshop

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.workshopService.GetOneWorkshop(input.ID)
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

func (h *workshopHandler) GetAllWorkshop(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.workshopService.GetAllWorkshop(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "Eror")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, newBerita)
	c.JSON(http.StatusOK, response)
}

func (h *workshopHandler) UpdateWorkshop(c *gin.Context) {
	// Parse the workshop ID from the URL parameter
	workshopID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.APIresponse(http.StatusBadRequest, "Invalid workshop ID")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Bind the input data from the request body to an UpdateWorkshop struct
	var input workshop.UpdateWorkshop
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Call the service to update the workshop
	updatedWorkshop, err := h.workshopService.UpdateWorkshop(workshopID, input)
	if err != nil {
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to update workshop")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, updatedWorkshop)
	c.JSON(http.StatusOK, response)
}

func (h *workshopHandler) DeleteWorkshop(c *gin.Context) {
	var input workshop.GetWorkshop

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDel, err := h.workshopService.DeleteWorkshop(input.ID)
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
