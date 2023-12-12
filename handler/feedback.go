package handler

import (
	"fmt"
	"greenwelfare/email"
	"greenwelfare/feedback"
	"greenwelfare/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type feedbackHandler struct {
	feedbackService feedback.Service
}

func NewFeedbackHandler(feedbackService feedback.Service) *feedbackHandler {
	return &feedbackHandler{feedbackService}
}

// @Summary Hapus feedback berdasarkan slug
// @Description Hapus feedback berdasarkan slug yang diberikan
// @Accept json
// @Produce json
// @Tags Feedback
// @Security BearerAuth
// @Param slug path int true "slug Feedback"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /feedback/{slug} [delete]

func (h *feedbackHandler) DeleteFeedback(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.feedbackService.DeleteFeedback(param)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	response := helper.APIresponse(http.StatusOK, "feedback has succesfuly deleted")
	c.JSON(http.StatusOK, response)

}

// @Summary Dapatkan feedback berdasarkan slug
// @Description Dapatkan feedback berdasarkan slug yang diberikan
// @Accept json
// @Produce json
// @Tags Feedback
// @Param slug path int true "slug Feedback"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /feedback/{slug} [get]
func (h *feedbackHandler) GetFeedbackBySlug(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.feedbackService.GetFeedbackBySlug(param)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, newDel)
	c.JSON(http.StatusOK, response)
}

// @Summary Dapatkan semua feedback atau feedback berdasarkan ID tertentu
// @Description Dapatkan semua feedback atau feedback berdasarkan ID tertentu
// @Accept json
// @Produce json
// @Tags Feedback
// @Param id query int false "ID Feedback"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /feedback [get]
func (h *feedbackHandler) GetAllFeedback(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	data, err := h.feedbackService.GetAllFeedback(input)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}

// @Summary Buat feedback baru
// @Description Buat feedback baru dengan informasi yang diberikan
// @Accept json
// @Produce json
// @Tags Feedback
// @Param Input body feedback.FeedbackInput true "Data Feedback yang ingin dibuat"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /feedback [post]
func (h *feedbackHandler) PostFeedbackHandler(c *gin.Context) {
	var feedbackInput feedback.FeedbackInput

	err := c.ShouldBind(&feedbackInput)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errorMessages}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newFeedbackPost, err := h.feedbackService.CreateFeedback(feedbackInput)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, "nil")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = email.SendEmailFeedback("raihanalfarisi2@gmail.com", feedbackInput.Email, feedbackInput.Text)
	if err != nil {
		// Handle kesalahan pengiriman email di sini.
		// Mungkin menampilkan pesan kesalahan kepada pengguna atau mencatatnya.
		fmt.Println("Error sending email:", err)
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, feedback.PostFormatterFeedback(newFeedbackPost))
	c.JSON(http.StatusOK, response)
}
