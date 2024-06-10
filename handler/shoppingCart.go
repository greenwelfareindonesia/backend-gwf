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

	ctx.JSON(http.StatusCreated, shoppingCart)
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

	ctx.JSON(http.StatusOK, shoppingCarts)
}
