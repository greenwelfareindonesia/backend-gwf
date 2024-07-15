package handler

import (
	"fmt"
	"greenwelfare/email"
	"greenwelfare/helper"
	"greenwelfare/dto"
	"greenwelfare/formatter"
	"greenwelfare/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type feedbackHandler struct {
	feedbackService service.ServiceFeedback
}

func NewFeedbackHandler(feedbackService service.ServiceFeedback) *feedbackHandler {
	return &feedbackHandler{feedbackService}
}

// @Summary Delete feedback slug
// @Description Delete feedback slug 
// @Accept json
// @Produce json
// @Tags Feedback
// @Security BearerAuth
// @Param slug path string true "slug Feedback"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/feedback/{slug} [delete]

func (h *feedbackHandler) DeleteFeedback(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.feedbackService.DeleteFeedback(param)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), param)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	response := helper.SuccessfulResponse1("feedback has succesfuly deleted")
	c.JSON(http.StatusOK, response)

}

// @Summary Get one feedback by slug
// @Description Get one feedback by slug 
// @Accept json
// @Produce json
// @Tags Feedback
// @Security BearerAuth
// @Param slug path string true "slug Feedback"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/feedback/{slug} [get]
func (h *feedbackHandler) GetFeedbackBySlug(c *gin.Context) {
	param := c.Param("slug")

	newDel, err := h.feedbackService.GetFeedbackBySlug(param)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), newDel)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(newDel)
	c.JSON(http.StatusOK, response)
}

// @Summary Get All feedback 
// @Description Get All feedback 
// @Accept json
// @Produce json
// @Tags Feedback
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/feedback [get]
func (h *feedbackHandler) GetAllFeedback(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	data, err := h.feedbackService.GetAllFeedback(input)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.SuccessfulResponse1(data)
	c.JSON(http.StatusOK, response)
}

// @Summary Create New feedback 
// @Description Create New feedback 
// @Accept json
// @Produce json
// @Tags Feedback
// @Param Input body dto.FeedbackInput true "Data for Create Feedback"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Router /api/feedback [post]
func (h *feedbackHandler) PostFeedbackHandler(c *gin.Context) {
	var feedbackInput dto.FeedbackInput

	err := c.ShouldBind(&feedbackInput)
	if err != nil {
		// errorMessages := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errorMessages}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), feedbackInput)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newFeedbackPost, err := h.feedbackService.CreateFeedback(feedbackInput)
	if err != nil {
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, "nil", newFeedbackPost)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = email.SendEmailFeedback("raihanalfarisi2@gmail.com", feedbackInput.Email, feedbackInput.Text)
	if err != nil {
		// Handle kesalahan pengiriman email di sini.
		// Mungkin menampilkan pesan kesalahan kepada pengguna atau mencatatnya.
		fmt.Println("Error sending email:", err)
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccessfulResponse1(formatter.PostFormatterFeedback(newFeedbackPost))
	c.JSON(http.StatusOK, response)
}
