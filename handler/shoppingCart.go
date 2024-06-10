package handler

import (
	"greenwelfare/dto"
	"greenwelfare/entity"
	"greenwelfare/helper"
	"greenwelfare/service"
	"net/http"

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
		ctx.JSON(http.StatusInternalServerError, response)
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
