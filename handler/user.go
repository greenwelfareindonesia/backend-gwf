package handler

import (
	"greenwelfare/auth"
	"greenwelfare/helper"
	"greenwelfare/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// @Summary Register new user
// @Description Register a new user with the provided information
// @Accept json
// @Produce json
// @Tags Users
// @Param body body user.RegisterUserInput true "User registration details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/user/register [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// errors := helper.FormatValidationError(err)
		// errorMessage := gin.H{"errors": errors}
		response := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), input)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Periksa ketersediaan email sebelum mendaftarkan pengguna
	isEmailAvailable, err := h.userService.IsEmaillAvailabilty(input.Email)
	if err != nil {
		response := helper.FailedResponse1(http.StatusConflict, err.Error(), input)
		c.JSON(http.StatusConflict, response)
		return
	}

	// Jika email tidak tersedia, kirim respons kesalahan
	if !isEmailAvailable {
		response := helper.FailedResponse1(http.StatusConflict, err.Error(), input)
		c.JSON(http.StatusConflict, response)
		return
	}

	// Register user jika email tersedia
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Format dan kirim respons berhasil jika registrasi berhasil
	response := helper.SuccessfulResponse1(newUser)
	c.JSON(http.StatusOK, response)
}

// @Summary User login
// @Description Log in an existing user using email and password
// @Accept json
// @Produce json
// @Tags Users
// @Param body body user.LoginInput true "User login credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/user/login [post]
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedinUser.ID, loggedinUser.Role)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.PostFormatterUser(loggedinUser, token)
	response := helper.APIresponse(http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

// @Summary Delete user
// @Description Delete a user
// @Security BearerAuth
// @Produce json
// @Tags Users
// @Param slug path string true "Slug User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/user/{slug} [delete]
func (h *userHandler) DeletedUser(c *gin.Context) {
	param := c.Param("slug")

	_, err := h.userService.DeleteUser(param)
	if err != nil {
		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// formatter := user.FormatterUser(newDel, "nil")
	response := helper.APIresponse(http.StatusOK, "Account has succesfuly deleted")
	c.JSON(http.StatusOK, response)
}

// @Summary Update user information
// @Description Update user details by ID
// @Security BearerAuth
// @Accept json
// @Produce json
// @Tags Users
// @Param slug path string true "User Slug"
// @Param body body user.UpdateUserInput true "User information for update"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 422 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/user/{slug} [put]
func (h *userHandler) UpdateUser(c *gin.Context) {
	param := c.Param("slug")
	var input user.UpdateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIresponse(http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.UpdateUser(param, input)
	if err != nil {

		response := helper.APIresponse(http.StatusUnprocessableEntity, err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIresponse(http.StatusOK, user)
	c.JSON(http.StatusOK, response)
}

//create folder
