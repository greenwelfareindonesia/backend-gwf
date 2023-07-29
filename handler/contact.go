package handler

import (
	"greenwelfare/contact"
	"greenwelfare/helper"
	"net/http"

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
