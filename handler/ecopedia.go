package handler

import (
	"fmt"
	"greenwelfare/ecopedia"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ecopediaHandler struct {
	ecopediaService ecopedia.Service
}

func NewEcopediaHandler(ecopediaService ecopedia.Service) *ecopediaHandler {
	return &ecopediaHandler{ecopediaService}
}

func (h *ecopediaHandler) GetEcopediasHandler(c *gin.Context) {
	ecopedias, err := h.ecopediaService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var ecopediasFormatter []ecopedia.EcopediaFormatter

	for _, e := range ecopedias {
		ecopediaFormatter := ecopedia.EcopediaFormatter{
			Judul:     e.Judul,
			Subjudul:  e.Subjudul,
			Deskripsi: e.Deskripsi,
			Gambar:    e.Gambar,
			Srcgambar: e.Srcgambar,
			Referensi: e.Referensi,
		}

		ecopediasFormatter = append(ecopediasFormatter, ecopediaFormatter)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ecopediasFormatter,
	})
}

func (h *ecopediaHandler) GetEcopediaHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	ecopedia, err := h.ecopediaService.FindById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ecopedia,
	})
}

func (h *ecopediaHandler) PostEcopediaHandler(c *gin.Context) {
	var ecopediaInput ecopedia.EcopediaInput

	err := c.ShouldBindJSON(&ecopediaInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ecopedia, err := h.ecopediaService.CreateEcopedia(ecopediaInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ecopedia,
	})
}
