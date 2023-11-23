package handler

import (
	"fmt"
	"greenwelfare/contact"
	"greenwelfare/email"
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

// @Summary Submit Contact Form
// @Description Submit a contact form
// @Accept json
// @Produce json
// @Tags Contact
// @Param requestBody body contact.ContactSubmissionInput true "Contact form input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /contact [post]
func (h *contactHandler) SubmitContactForm(c *gin.Context) {
	var input contact.ContactSubmissionInput
	

	err := c.ShouldBindJSON(&input)
	// fmt.Println(err)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newContactSubmission, err := h.contactService.SubmitContactSubmission(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "nil")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// emailBody := "Terima kasih atas pesan Anda. Kami akan segera menghubungi Anda."
	err = email.SendEmail("raihanalfarisi2@gmail.com", input.Subject, input.Name, input.Email, input.Message)
	if err != nil {
    // Handle kesalahan pengiriman email di sini.
    // Mungkin menampilkan pesan kesalahan kepada pengguna atau mencatatnya.
	fmt.Println("Error sending email:", err)
		response := helper.APIresponse(http.StatusUnprocessableEntity, "nilll")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	formatter := contact.FormatterContact(newContactSubmission)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

// @Summary Get All Contact Submissions
// @Description Get all contact form submissions
// @Accept json
// @Produce json
// @Tags Contact
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /contact [get]
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

// @Summary Get Contact Submission by ID
// @Description Get a contact form submission by ID
// @Accept json
// @Produce json
// @Tags Contact
// @Param id path int true "Contact Submission ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /contact/{id} [get]
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

// @Summary Delete Contact Submission by ID
// @Description Delete a contact form submission by ID
// @Accept json
// @Produce json
// @Tags Contact
// @Param id path int true "Contact Submission ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /contact/{id} [delete]
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
