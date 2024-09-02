package handler

import (
	"greenwelfare/dto"
	"greenwelfare/helper"
	"greenwelfare/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kajianHandler struct {
	kajianService service.ServiceKajian
}

func NewKajianHandler() *kajianHandler {
	return &kajianHandler{}
}

func (h *kajianHandler) CreateKajian(c *gin.Context) {
	// TODO: upload images

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
