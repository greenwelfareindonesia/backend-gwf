package handler

import (
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/helper"
	"greenwelfare/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type shoppingCartHandler struct {
	shoppingCartService service.ServiceShoppingCart
}

func NewShoppingCartHandler(svc service.ServiceShoppingCart) *shoppingCartHandler {
	return &shoppingCartHandler{svc}
}

func (h *shoppingCartHandler) CreateShoppingCart(ctx *gin.Context) {
	newShoppingCartRequest := dto.CreateShoppingCartDTO{}

	// BINDING
	if err := ctx.Bind(&newShoppingCartRequest); err != nil {
		errBindJson := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), newShoppingCartRequest)
		ctx.AbortWithStatusJSON(errBindJson.Error.Code, errBindJson)
		return
	}

	// VALIDATE INPUT
	if err := newShoppingCartRequest.Validate(); err != nil {
		errValidate := helper.FailedResponse1(http.StatusBadRequest, err.Error(), newShoppingCartRequest)
		ctx.AbortWithStatusJSON(errValidate.Error.Code, errValidate)
		return
	}

	// SET USER_ID BY SESSION
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		response := helper.FailedResponse1(http.StatusInternalServerError, "User Session not found", nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	newShoppingCartRequest.UserID = uint64(currentUser.(*entity.User).ID)

	// SAVE SHOPPING CART
	shoppingCart, errSvc := h.shoppingCartService.CreateShoppingCart(ctx, newShoppingCartRequest)
	if errSvc != nil {
		errCreate := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), newShoppingCartRequest)
		ctx.AbortWithStatusJSON(errCreate.Error.Code, errCreate)
		return
	}

	response := helper.SuccessfulResponse1(shoppingCart)
	ctx.JSON(http.StatusOK, response)
}

func (h *shoppingCartHandler) GetShoppingCarts(ctx *gin.Context) {
	userID := uint64(0)

	// FILTER USER_ID BY SESSION USER
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		response := helper.FailedResponse1(http.StatusInternalServerError, "User Session not found", nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	roleId := uint64(currentUser.(*entity.User).Role)
	if roleId == 0 {
		userID = uint64(currentUser.(*entity.User).ID)
	}

	// FILTER USER_ID BY QUERY PARAM (ONLY ADMIN)
	queryUserID := ctx.Query("user_id")
	if queryUserID != "" {
		id, err := strconv.Atoi(queryUserID)
		if id == 0 || err != nil {
			response := helper.FailedResponse1(http.StatusBadRequest, "invalid query param", nil)
			ctx.AbortWithStatusJSON(response.Error.Code, response)
			return
		}

		// JIKA ROLE ADMIN MAKA BOLEH MELAKUKAN FILTER BERDASARKAN USER_ID MELALUI QUERY PARAM
		if roleId == 1 {
			userID = uint64(id)
		}
	}

	shoppingCarts, errSvc := h.shoppingCartService.GetShoppingCarts(ctx, userID)
	if errSvc != nil {
		errGetShoppingCarts := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), errSvc)
		ctx.AbortWithStatusJSON(errGetShoppingCarts.Error.Code, errGetShoppingCarts)
		return
	}

	response := helper.SuccessfulResponse1(shoppingCarts)
	ctx.JSON(http.StatusOK, response)
}

func (h *shoppingCartHandler) GetShoppingCartById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, "invalid path param id", nil)
		ctx.AbortWithStatusJSON(response.Error.Code, response)
		return
	}

	userID := uint64(0)

	// JIKA LOGIN SBG USER . WAJIT MEMBERIKAN PARAMETR USER_ID AGAR HANYA DATANYA SENDIRI YG BISA DILIHAT
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		response := helper.FailedResponse1(http.StatusInternalServerError, "user Session not found", nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	// SET USER_ID
	roleId := uint64(currentUser.(*entity.User).Role)
	if roleId == 0 {
		userID = uint64(currentUser.(*entity.User).ID)
	}

	shoppingCart, errSvc := h.shoppingCartService.GetShoppingCartById(ctx, userID, uint64(id))
	if errSvc != nil {
		errGetShoppingCarts := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), errSvc)
		ctx.AbortWithStatusJSON(errGetShoppingCarts.Error.Code, errGetShoppingCarts)
		return
	}
	response := helper.SuccessfulResponse1(shoppingCart)
	ctx.JSON(http.StatusOK, response)
}

func (h *shoppingCartHandler) UpdateShoppingCartById(ctx *gin.Context) {
	updateShoppingCart := dto.UpdateShoppingCartDTO{}

	// GET PARAMS ID
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, "invalid path param id", nil)
		ctx.AbortWithStatusJSON(response.Error.Code, response)
		return
	}

	// BINDING
	if err := ctx.Bind(&updateShoppingCart); err != nil {
		errBindJson := helper.FailedResponse1(http.StatusUnprocessableEntity, err.Error(), updateShoppingCart)
		ctx.AbortWithStatusJSON(errBindJson.Error.Code, errBindJson)
		return
	}

	// VALIDATE INPUT
	if err := updateShoppingCart.Validate(); err != nil {
		errValidate := helper.FailedResponse1(http.StatusBadRequest, err.Error(), updateShoppingCart)
		ctx.AbortWithStatusJSON(errValidate.Error.Code, errValidate)
		return
	}

	userID := uint64(0)

	// JIKA LOGIN SBG USER . WAJIT MEMBERIKAN PARAMETR USER_ID AGAR HANYA DATANYA SENDIRI YG BISA DILIHAT
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		response := helper.FailedResponse1(http.StatusInternalServerError, "user Session not found", nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	// SET USER_ID
	roleId := uint64(currentUser.(*entity.User).Role)
	if roleId == 0 {
		userID = uint64(currentUser.(*entity.User).ID)
	}

	// UPDATE DTO
	updateShoppingCart.ID = uint64(id)
	updateShoppingCart.UserID = userID

	shoppingCart, errSvc := h.shoppingCartService.UpdateShoppingCartById(ctx, updateShoppingCart)
	if errSvc != nil {
		errUpdateShoppingCarts := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), errSvc)
		ctx.AbortWithStatusJSON(errUpdateShoppingCarts.Error.Code, errUpdateShoppingCarts)
		return
	}
	response := helper.SuccessfulResponse1(shoppingCart)
	ctx.JSON(http.StatusOK, response)
}

func (h *shoppingCartHandler) DeleteShoppingCartById(ctx *gin.Context) {
	// GET PARAMS ID
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		response := helper.FailedResponse1(http.StatusBadRequest, "invalid path param id", nil)
		ctx.AbortWithStatusJSON(response.Error.Code, response)
		return
	}

	userID := uint64(0)

	// JIKA LOGIN SBG USER . WAJIT MEMBERIKAN PARAMETR USER_ID AGAR HANYA DATANYA SENDIRI
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		response := helper.FailedResponse1(http.StatusInternalServerError, "user Session not found", nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	// SET USER_ID
	roleId := uint64(currentUser.(*entity.User).Role)
	if roleId == 0 {
		userID = uint64(currentUser.(*entity.User).ID)
	}

	// DELETE SHOPPING CART
	errSvc := h.shoppingCartService.DeleteShoppingCartById(ctx, userID, uint64(id))
	if errSvc != nil {
		errDeleteShoppingCart := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), errSvc)
		ctx.AbortWithStatusJSON(errDeleteShoppingCart.Error.Code, errDeleteShoppingCart)
		return
	}

	response := helper.SuccessfulResponse1("success delete shopping cart")
	ctx.JSON(http.StatusOK, response)
}

func (h *shoppingCartHandler) GetStatisticCarts(ctx *gin.Context) {
	userID := uint64(0)

	// FILTER USER_ID BY SESSION USER
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		response := helper.FailedResponse1(http.StatusInternalServerError, "User Session not found", nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	roleId := uint64(currentUser.(*entity.User).Role)
	if roleId == 0 {
		userID = uint64(currentUser.(*entity.User).ID)
	}

	// FILTER USER_ID BY QUERY PARAM (ONLY ADMIN)
	queryUserID := ctx.Query("user_id")
	if queryUserID != "" {
		id, err := strconv.Atoi(queryUserID)
		if id == 0 || err != nil {
			response := helper.FailedResponse1(http.StatusBadRequest, "invalid query param", nil)
			ctx.AbortWithStatusJSON(response.Error.Code, response)
			return
		}

		// JIKA ROLE ADMIN MAKA BOLEH MELAKUKAN FILTER BERDASARKAN USER_ID MELALUI QUERY PARAM
		if roleId == 1 {
			userID = uint64(id)
		}
	}

	statisticData, errSvc := h.shoppingCartService.GetStatisticCarts(ctx, userID)
	if errSvc != nil {
		errGetStatisticCarts := helper.FailedResponse1(http.StatusInternalServerError, errSvc.Error(), errSvc)
		ctx.AbortWithStatusJSON(errGetStatisticCarts.Error.Code, errGetStatisticCarts)
		return
	}

	response := helper.SuccessfulResponse1(statisticData)
	ctx.JSON(http.StatusOK, response)
}
