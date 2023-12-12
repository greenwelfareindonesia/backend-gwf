package handler

import (
	"bytes"
	"context"
	"fmt"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/event"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	eventService event.Service
	endpointService endpointcount.StatisticsService
}

func NewEventHandler(eventService event.Service, endpointService endpointcount.StatisticsService) *eventHandler {
	return &eventHandler{eventService, endpointService}
}

// @Summary Hapus event berdasarkan ID
// @Description Hapus event berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Tags Event
// @Security BearerAuth
// @Param id path int true "ID Event"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /event/{id} [delete]
func (h *eventHandler) DeleteEvent(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.eventService.DeleteEvent(param)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.APIresponse(http.StatusOK, "event has succesfuly deleted")
	c.JSON(http.StatusOK, response)

}

// @Summary Dapatkan satu event berdasarkan ID
// @Description Dapatkan satu event berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Tags Event
// @Param id path int true "ID Event"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /event/{id} [get]
func (h *eventHandler) GetOneEvent(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.eventService.GetOneEvent(param)
	if err != nil {
		
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetByIDEvent /Event/GetByIDEvent", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusOK, newDel)
	c.JSON(http.StatusOK, response)

}

// @Summary Update event berdasarkan ID
// @Description Update event berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Tags Event
// @Security BearerAuth
// @Param id path int true "ID Event"
// @Param File formData file true "File gambar"
// @Param Judul formData string true "Judul"
// @Param Message formData string true "Pesan Event"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /event/{id} [put]
func (h *eventHandler) UpdateEvent(c *gin.Context) {
	file, _ := c.FormFile("File")
	src, err := file.Open()
	if err != nil {
		fmt.Printf("error when open file %v", err)
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v", err)
		return
	}

	img, err := imagekits.Base64toEncode(buf.Bytes())
	if err != nil {
		fmt.Println("error reading image ", err)
	}

	fmt.Println("image base 64 format : ", img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	param := c.Param("slug")

	var input event.CreateEvents
	err = c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	events, err := h.eventService.UpdateEvent(param,input,imageKitURL)
	if err != nil {
		
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}


	response := helper.APIresponse(http.StatusOK, event.UpdatedFormatterEvent(events))
	c.JSON(http.StatusOK, response)

}

// @Summary Buat event baru
// @Description Buat event baru dengan informasi yang diberikan
// @Accept json
// @Produce json
// @Tags Event
// @Security BearerAuth
// @Param File formData file true "File gambar"
// @Param Judul formData string true "Judul"
// @Param Message formData string true "Pesan Event"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /event [post]
func (h *eventHandler) CreateEvent(c *gin.Context) {
	file, _ := c.FormFile("File")
	src, err := file.Open()
	if err != nil {
		fmt.Printf("error when open file %v", err)
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		fmt.Printf("error read file %v", err)
		return
	}

	img, err := imagekits.Base64toEncode(buf.Bytes())
	if err != nil {
		fmt.Println("error reading image ", err)
	}

	fmt.Println("image base 64 format : ", img)

	imageKitURL, err := imagekits.ImageKit(context.Background(), img)
	if err != nil {
		// Tangani jika terjadi kesalahan saat upload gambar
		// Misalnya, Anda dapat mengembalikan respon error ke klien jika diperlukan
		response := helper.APIresponse(http.StatusInternalServerError, "Failed to upload image")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	var input event.CreateEvents

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
		
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.eventService.CreateEvent(input, imageKitURL)
	if err != nil {
		// data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, event.PostFormatterEvent(data))
	c.JSON(http.StatusOK, response)
}

// @Summary Dapatkan semua event atau event berdasarkan ID tertentu
// @Description Dapatkan semua event atau event berdasarkan ID tertentu
// @Accept json
// @Produce json
// @Tags Event
// @Param id query int false "ID Event"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /event [get]
func (h *eventHandler) GetAllEvent(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.eventService.GetAllEvent(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userAgent := c.GetHeader("User-Agent")

	err = h.endpointService.IncrementCount("GetAllEvent /Event/GetAllEvent", userAgent)
    if err != nil {
        response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
    }

	response := helper.APIresponse(http.StatusOK, newBerita)
	c.JSON(http.StatusOK, response)
}
