package handler

import (
	"greenwelfare/dto"
	"greenwelfare/helper"
	"greenwelfare/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type hrdHandler struct {
	hrdService service.ServiceHrd
}

func NewHrdHandler(hrdService service.ServiceHrd) *hrdHandler {
	return &hrdHandler{hrdService}
}

func (h *hrdHandler) CreateHrd(c *gin.Context) {
	var input dto.CreateHrdDTO
	if err := c.ShouldBind(&input); err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.hrdService.CreateHrd(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusCreated, response)
}

func (h *hrdHandler) GetOneHrd(c *gin.Context) {
	slug := c.Param("slug")

	data, err := h.hrdService.GetOneHrd(slug)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)
}

func (h *hrdHandler) GetAllHrd(c *gin.Context) {
	data, err := h.hrdService.GetAllHrd()
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)
}

func (h *hrdHandler) UpdateHrd(c *gin.Context) {
	slug := c.Param("slug")

	var input dto.UpdateHrdDTO
	if err := c.ShouldBind(&input); err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.hrdService.UpdateHrd(input, slug)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)
}

func (h *hrdHandler) DeleteHrd(c *gin.Context) {
	slug := c.Param("slug")

	if _, err := h.hrdService.DeleteHrd(slug); err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1("hrd has successfully deleted")
	c.JSON(http.StatusOK, response)
}
