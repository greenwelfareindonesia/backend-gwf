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


// @Summary Hapus data Ecopedia berdasarkan ID
// @Description Hapus data Ecopedia berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Ecopedia ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ecopedia/{id} [delete]
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

// @Summary Dapatkan data Ecopedia berdasarkan ID
// @Description Dapatkan data Ecopedia berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Param id path int true "Ecopedia ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ecopedia/{id} [get]
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

// @Summary Dapatkan semua data Ecopedia
// @Description Dapatkan semua data Ecopedia dengan opsi ID sebagai parameter query opsional
// @Accept json
// @Produce json
// @Param id query int false "ID Ecopedia (opsional)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ecopedia [get]
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

// @Summary Update data Ecopedia
// @Description Update data Ecopedia berdasarkan ID yang diberikan dengan informasi yang diberikan
// @Accept multipart/form-data
// @Security BearerAuth
// @Produce json
// @Param id path int true "Ecopedia ID"
// @Param file formData file true "File gambar"
// @Param judul formData string true "judul"
// @Param SubJudul formData string true "sub_judul"
// @Param Deskripsi formData string true "deskripsi"
// @Param SrcGambar formData string true "src_gambar"
// @Param Referensi formData string true "referensi"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ecopedia/{id} [put]
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
	if err!=nil{
		fmt.Printf("error when open file %v",err)
	}
	defer	src.Close()

	
	buf:=bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v",err)
		return 
	}	

	img,err:=imagekits.Base64toEncode(buf.Bytes())
	if err!=nil{
		fmt.Println("error reading image ",err)
	}

	fmt.Println("image base 64 format : ",img)

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

// @Summary Buat data Ecopedia baru
// @Description Buat data Ecopedia baru dengan informasi yang diberikan
// @Accept multipart/form-data
// @Security BearerAuth
// @Produce json
// @Param file formData file true "File gambar"
// @Param judul formData string true "judul"
// @Param SubJudul formData string true "sub_judul"
// @Param Deskripsi formData string true "deskripsi"
// @Param SrcGambar formData string true "src_gambar"
// @Param Referensi formData string true "referensi"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ecopedia [post]
func (h *ecopediaHandler) PostEcopediaHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	src,err:=file.Open()
	if err!=nil{
		fmt.Printf("error when open file %v",err)
	}
	defer	src.Close()

	
	
	buf:=bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v",err)
		return 
	}	

	img,err:=imagekits.Base64toEncode(buf.Bytes())
	if err!=nil{
		fmt.Println("error reading image", err)
	}

	fmt.Println("image base 64 format",img)

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


// @Summary Buat komentar pada data Ecopedia
// @Description Buat komentar pada data Ecopedia berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Ecopedia ID"
// @Param body body string true "Komentar"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ecopedia/{id}/comment [post]
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
