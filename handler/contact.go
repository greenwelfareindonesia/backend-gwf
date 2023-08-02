package handler

import (
	"greenwelfare/contact"
	"greenwelfare/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type contactHandler struct {
	contactService contact.Service
}

func NewContactHandler(contactService contact.Service) *contactHandler {
	return &contactHandler{contactService}
}

func (h *contactHandler) SubmitContactForm(c *gin.Context) {
	var input contact.ContactSubmissionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newContactSubmission, err := h.contactService.SubmitContactSubmission(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := contact.FormatterContact(newContactSubmission)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *contactHandler) GetContactSubmissionsHandler(c *gin.Context) {
	contact_submissions, err := h.contactService.GetAllContactSubmission()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": contact_submissions,
	})
}

func (h *contactHandler) GetContactSubmissionHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	contact_submission, err := h.contactService.GetContactSubmissionById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": contact_submission,
	})
}

func (h *contactHandler) DeleteContactSubmissionHandler(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	h.contactService.DeleteContactSubmission(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete contact submission",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contact submission deleted",
	})
}
