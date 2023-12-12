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

// @Summary Hapus feedback berdasarkan ID
// @Description Hapus feedback berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Tags Feedback
// @Security BearerAuth
// @Param id path int true "ID Feedback"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /feedback/{id} [delete]

// func (h *feedbackHandler) DeleteFeedback(c *gin.Context) {
// 	var input feedback.FeedbackID

// 	err := c.ShouldBindUri(&input)

// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorMessage := gin.H{"errors": errors}
// 		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	data, err := h.feedbackService.DeleteFeedback(input.ID)
// 	if err != nil {

// 		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}
// 	response := helper.APIresponse(http.StatusOK, (data))
// 	c.JSON(http.StatusOK, response)
// }

func (h *feedbackHandler) DeleteFeedback(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.feedbackService.DeleteFeedback(param)
	if err != nil {
<<<<<<< HEAD
=======
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.feedbackService.DeleteFeedback(input.ID)
	if err != nil {
		
>>>>>>> 0075203b2f39f41648b074b4c80756123471ee58
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
<<<<<<< Updated upstream
	response := helper.APIresponse(http.StatusOK, "feedback has succesfuly deleted")
=======
<<<<<<< HEAD
	response := helper.APIresponse(http.StatusOK, newDel)
=======
	response := helper.APIresponse(http.StatusOK, "feedback has succesfuly deleted")
>>>>>>> 0075203b2f39f41648b074b4c80756123471ee58
>>>>>>> Stashed changes
	c.JSON(http.StatusOK, response)

}

// @Summary Dapatkan feedback berdasarkan ID
// @Description Dapatkan feedback berdasarkan ID yang diberikan
// @Accept json
// @Produce json
// @Tags Feedback
// @Param id path int true "ID Feedback"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /feedback/{id} [get]
func (h *feedbackHandler) GetFeedbackBySlug(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.feedbackService.GetFeedbackBySlug(param)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
<<<<<<< Updated upstream
=======
<<<<<<< HEAD
=======
>>>>>>> Stashed changes
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
>>>>>>> 0075203b2f39f41648b074b4c80756123471ee58

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
