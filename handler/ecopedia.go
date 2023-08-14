package handler

import (
	"bytes"
	"context"
	"fmt"
	"greenwelfare/ecopedia"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ecopediaHandler struct {
	ecopediaService ecopedia.Service
}

func NewEcopediaHandler(ecopediaService ecopedia.Service) *ecopediaHandler {
	return &ecopediaHandler{ecopediaService}
}

func (h *ecopediaHandler) GetAllEcopedia(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	ecopedia, err := h.ecopediaService.GetAllEcopedia(input)
	if ecopedia != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, ecopedia)
	c.JSON(http.StatusOK, response)
}

// func (h *ecopediaHandler) GetEcopediaHandler(c *gin.Context) {
// 	idString := c.Param("id")
// 	id, _ := strconv.Atoi(idString)

// 	ecopedia, err := h.ecopediaService.FindById(id)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": ecopedia,
// 	})
// }

func (h *ecopediaHandler) PostEcopediaHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	src,err:=file.Open()
	defer	src.Close()

	if err!=nil{
		fmt.Printf("error when open file %v",err)
	}
	
	buf:=bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v",err)
		return 
	}	

	img,err:=imagekits.Base64toEncode(buf.Bytes())
	if err!=nil{
		fmt.Println("error reading image %v",err)
	}

	fmt.Println("image base 64 format : %v",img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var ecopediaInput ecopedia.EcopediaInput

	err = c.ShouldBind(&ecopediaInput)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errorMessages}
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
		

	_, err = h.ecopediaService.CreateEcopedia(ecopediaInput, imageKitURL)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
