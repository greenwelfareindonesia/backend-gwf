package handler

import (
	"bytes"
	"context"
	"fmt"
	"greenwelfare/dto"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/formatter"
	"greenwelfare/helper"
	"greenwelfare/imagekits"
	"greenwelfare/service"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type eventHandler struct {
	eventService    service.ServiceEvent
	endpointService endpointcount.StatisticsService
}

func NewEventHandler(eventService service.ServiceEvent, endpointService endpointcount.StatisticsService) *eventHandler {
	return &eventHandler{eventService, endpointService}
}

// @Summary Delete event by slug
// @Description Delete event by slug
// @Accept json
// @Produce json
// @Tags Event
// @Security BearerAuth
// @Param slug path string true "slug Event"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/event/{slug} [delete]
func (h *eventHandler) DeleteEvent(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.eventService.DeleteEvent(param)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	response := helper.SuccessfulResponse1("event has succesfuly deleted")
	c.JSON(http.StatusOK, response)

}

// @Summary Get One Event
// @Description Get One Event by slug
// @Accept json
// @Produce json
// @Tags Event
// @Param slug path string true "slug event"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/event/{slug} [get]
func (h *eventHandler) GetOneEvent(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.eventService.GetOneEvent(param)
	if err != nil {

		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), newDel)
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

	response := helper.SuccessfulResponse1(newDel)
	c.JSON(http.StatusOK, response)

}

// @Summary Update event by slug
// @Description Update event by slug
// @Accept json
// @Produce json
// @Tags Event
// @Security BearerAuth
// @Param slug path string true "slug event"
// @Param file formData file true "File Image"
// @Param title formData string true "Title Event"
// @Param eventMessage formData string true "Event Message"
// @Param location formData string true "Location"
// @Param date formData string true "Date"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/event/{slug} [put]
func (h *eventHandler) UpdateEvent(c *gin.Context) {
	file, _ := c.FormFile("file")
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

	var input dto.CreateEvents
	err = c.ShouldBind(&input)
	if err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	events, err := h.eventService.UpdateEvent(param, input, imageKitURL)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(formatter.UpdatedFormatterEvent(events))
	c.JSON(http.StatusOK, response)

}

// @Summary Create New Event
// @Description Create New Event
// @Accept json
// @Produce json
// @Tags Event
// @Security BearerAuth
// @Param file formData file true "File Image"
// @Param title formData string true "Title Event"
// @Param eventMessage formData string true "Event Message"
// @Param location formData string true "Location"
// @Param date formData string true "Date"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/event [post]
func (h *eventHandler) CreateEvent(c *gin.Context) {
	file, _ := c.FormFile("file")
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

	var input dto.CreateEvents

	err = c.ShouldBind(&input)

	if err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
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
	response := helper.SuccessfulResponse1(formatter.PostFormatterEvent(data))
	c.JSON(http.StatusOK, response)
}

// @Summary Get All Event
// @Description Get All Event
// @Accept json
// @Produce json
// @Tags Event
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/event [get]
func (h *eventHandler) GetAllEvent(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	newBerita, err := h.eventService.GetAllEvent(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), newBerita)
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

	response := helper.SuccessfulResponse1(newBerita)
	c.JSON(http.StatusOK, response)
}
