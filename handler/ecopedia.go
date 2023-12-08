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
// @Tags Ecopedia
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
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
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
// @Tags Ecopedia
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
		
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
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
	
	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)

}

// @Summary Dapatkan semua data Ecopedia
// @Description Dapatkan semua data Ecopedia dengan opsi ID sebagai parameter query opsional
// @Accept json
// @Produce json
// @Tags Ecopedia
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
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
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

	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)
}

// @Summary Update data Ecopedia
// @Description Update data Ecopedia berdasarkan ID yang diberikan dengan informasi yang diberikan
// @Accept multipart/form-data
// @Security BearerAuth
// @Produce json
// @Tags Ecopedia
// @Param id path int true "Ecopedia ID"
// @Param file formData file true "File gambar"
// @Param judul formData string true "judul"
// @Param SubJudul formData string true "sub_judul"
// @Param Deskripsi formData string true "Deskripsi"
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
			response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		

	_, err = h.ecopediaService.UpdateEcopedia(inputID, input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
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
// @Tags Ecopedia
// @Param file1 formData file true "File gambar 1"
// @Param file2 formData file true "File gambar 2"
// @Param judul formData string true "judul"
// @Param SubJudul formData string true "sub_judul"
// @Param Deskripsi formData string true "Deskripsi"
// @Param SrcGambar formData string true "src_gambar"
// @Param Referensi formData string true "referensi"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ecopedia [post]
func (h *ecopediaHandler) PostEcopediaHandler(c *gin.Context) {
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

		var ecopediaInput ecopedia.EcopediaInput

		err := c.ShouldBind(&ecopediaInput)

		if err != nil {
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		// Create a new news item with the provided input
		newNews, err := h.ecopediaService.CreateEcopedia(ecopediaInput)
		fmt.Println(newNews)
		if err != nil {
			response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
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

// Respond with a success message
			data := gin.H{"is_uploaded": true}
			response := helper.APIresponse(http.StatusOK, data)
			c.JSON(http.StatusOK, response)

}


// @Summary Tambahkan komentar atau tindakan pengguna terhadap Ecopedia
// @Description Tambahkan komentar atau tindakan pengguna terhadap Ecopedia berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Tags Ecopedia
// @Param id path int true "ID Ecopedia"
// @Param comment formData string true "Komentar"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /ecopedia/comment/{id} [post]
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
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
