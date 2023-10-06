package handler

import (
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

func (h *feedbackHandler) DeleteFeedback(c *gin.Context) {
	var input feedback.FeedbackID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.feedbackService.DeleteFeedback(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)
}

func (h *feedbackHandler) GetFeedbackByID(c *gin.Context) {
	var input feedback.FeedbackID

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.feedbackService.GetFeedbackByID(input.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)

}

func (h *feedbackHandler) GetAllFeedback(c *gin.Context) {
	input, _ := strconv.Atoi(c.Query("id"))

	data, err := h.feedbackService.GetAllFeedback(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIresponse(http.StatusOK, (data))
	c.JSON(http.StatusOK, response)
}

func (h *feedbackHandler) UpdateFeedback(c *gin.Context) {
	var input feedback.FeedbackInput

	var inputID feedback.FeedbackID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBind(&input)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errorMessages}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		//inisiasi data yang tujuan dalam return hasil ke postman
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// _, err = h.feedbackService.UpdateFeedback(inputID, input)
	// if err != nil {
	// 	data := gin.H{"is_uploaded": false}
	// 	response := helper.APIresponse(http.StatusUnprocessableEntity, data)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	// data := gin.H{"is_uploaded": true}
	// response := helper.APIresponse(http.StatusOK, data)
	// c.JSON(http.StatusOK, response)

}

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

	_, err = h.feedbackService.CreateFeedback(feedbackInput)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse(http.StatusUnprocessableEntity, data)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse(http.StatusOK, data)
	c.JSON(http.StatusOK, response)
}
