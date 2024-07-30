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

// @Summary Create a new Hrd staff member
// @Description Register a new Hrd staff with the provided information
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Hrd
// @Param body body dto.CreateHrdDTO true "Hrd details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/hrd [post]
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

// @Summary Get one Hrd
// @Description Get one Hrd by slug
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Hrd
// @Param slug path string true "slug"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/hrd/{slug} [get]
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

// @Summary Get all Hrd
// @Description Get all Hrd
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Hrd
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/hrd [get]
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

// @Summary Get all Hrd by departement
// @Description Get all Hrd by departement
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Hrd
// @Param departement path string true "departement"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/hrd/departement/{departement} [get]
func (h *hrdHandler) GetAllByDepartement(c *gin.Context) {
	departement := c.Param("departement")

	data, err := h.hrdService.GetAllHrdByDepartement(departement)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)
}

// @Summary Get all Hrd by status
// @Description Get all Hrd by status
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Hrd
// @Param status path string true "status"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/hrd/status/{status} [get]
func (h *hrdHandler) GetAllByStatus(c *gin.Context) {
	status := c.Param("status")

	data, err := h.hrdService.GetAllHrdByStatus(status)
	if err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, err.Error(), err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)
}

// UpdateHrd godoc
// @Summary Update Hrd
// @Description Update Hrd
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Hrd
// @Param slug path string true "slug hrd"
// @Param body body dto.UpdateHrdDTO true "Hrd details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/hrd/{slug} [put]
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

// @Summary Delete Hrd
// @Description Delete hrd by slug
// @Accept json
// @Produce json
// @Security BearerAuth
// @Tags Hrd
// @Param slug path string true "slug"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/hrd/{slug} [delete]
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
