package handler

import (
	"greenwelfare/email"
	"greenwelfare/helper"
	"greenwelfare/dto"
	"greenwelfare/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type contactHandler struct {
	contactService service.ServiceContact
}

func NewContactHandler(contactService service.ServiceContact) *contactHandler {
	return &contactHandler{contactService}
}

// @Summary Submit Contact Form
// @Description Submit a contact form
// @Accept json
// @Produce json
// @Tags Contact
// @Param requestBody body dto.ContactSubmissionInput true "Contact form input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/contact [post]
func (h *contactHandler) SubmitContactForm(c *gin.Context) {
	var input dto.ContactSubmissionInput

	err := c.ShouldBindJSON(&input)
	// fmt.Println(err)
	if err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newContactSubmission, err := h.contactService.SubmitContactSubmission(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// emailBody := "Terima kasih atas pesan Anda. Kami akan segera menghubungi Anda."
	err = email.SendEmail("raihanalfarisi2@gmail.com", input.Subject, input.Name, input.Email, input.Message)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(newContactSubmission)
	c.JSON(http.StatusOK, response)
}

// @Summary Get All Contact Submissions
// @Description Get all contact form submissions
// @Accept json
// @Produce json
// @Tags Contact
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/contact [get]
func (h *contactHandler) GetContactSubmissionsHandler(c *gin.Context) {
	contact_submissions, err := h.contactService.GetAllContactSubmission()
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), contact_submissions)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(contact_submissions)
	c.JSON(http.StatusOK, response)
}

// @Summary Get Contact Submission by slug
// @Description Get a contact form submission by slug
// @Accept json
// @Produce json
// @Tags Contact
// @Security BearerAuth
// @Param slug path string true "Contact Submission slug"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/contact/{slug} [get]
func (h *contactHandler) GetContactSubmissionHandler(c *gin.Context) {
	param := c.Param("slug")

	contact_submission, err := h.contactService.GetContactSubmissionById(param)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(contact_submission)
	c.JSON(http.StatusOK, response)
}

// @Summary Delete Contact Submission by slug
// @Description Delete a contact form submission by slug
// @Accept json
// @Produce json
// @Tags Contact
// @Security BearerAuth
// @Param slug path string true "Contact Submission slug"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/contact/{slug} [delete]
func (h *contactHandler) DeleteContactSubmissionHandler(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.contactService.DeleteContactSubmission(param)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1("Contact submission deleted")
	c.JSON(http.StatusOK, response)
}
