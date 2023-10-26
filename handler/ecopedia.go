package handler

import (
	"bytes"
	"context"
	"fmt"
	"greenwelfare/ecopedia"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"greenwelfare/user"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ecopediaHandler struct {
	ecopediaService ecopedia.Service
	endpointService endpointcount.StatisticsService
}

func NewEcopediaHandler(ecopediaService ecopedia.Service, endpointService endpointcount.StatisticsService) *ecopediaHandler {
	return &ecopediaHandler{ecopediaService, endpointService}
}


func (h *ecopediaHandler) DeleteEcopedia (c *gin.Context){
	var input ecopedia.EcopediaID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.ecopediaService.DeleteEcopedia(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, ecopedia.FormatterEcopedia(data))
	c.JSON(http.StatusOK, response)
}

func (h *ecopediaHandler) GetEcopediaByID (c *gin.Context){
	var input ecopedia.EcopediaID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.ecopediaService.GetEcopediaByID(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDEcopedia /Ecopedia/GetByIDEcopedia", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)

}

func (h *ecopediaHandler) GetAllEcopedia (c *gin.Context){
	input, _ := strconv.Atoi(c.Query("id"))

	data, err := h.ecopediaService.GetAllEcopedia(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetAllEcopedia /Ecopedia/GetAllEcopedia", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)
}

func (h *ecopediaHandler) UpdateEcopedia (c *gin.Context) {
	var input ecopedia.EcopediaInput

	var inputID ecopedia.EcopediaID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

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
			data := gin.H{"is_uploaded": false}
			response := helper.APIresponse(http.StatusUnprocessableEntity, data)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		

	_, err = h.ecopediaService.UpdateEcopedia(inputID, input, imageKitURL)
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


func (h *ecopediaHandler) PostCommentEcopedia(c *gin.Context) {
	var getIdEcopedia ecopedia.EcopediaID

	err := c.ShouldBindUri(&getIdEcopedia)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errorMessages}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
		}

	var ecopediaInput ecopedia.UserActionToEcopedia

	currentUser := c.MustGet("currentUser").(user.User)
	// userId := currentUser.ID
	ecopediaInput.User = currentUser

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
		

	_, err = h.ecopediaService.UserActionToEcopedia(getIdEcopedia ,ecopediaInput)
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
