package handler

import (
	"go-technical-test-synapsis/src/cart"
	"go-technical-test-synapsis/src/entity"
	"go-technical-test-synapsis/src/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	cartService cart.Service
}

func NewCart(service cart.Service) *cartHandler {
	return &cartHandler{service}
}

func (h *cartHandler) Mount(cart *gin.RouterGroup) {
	cart.POST("/cart", h.StoreItemToCart)
	cart.POST("/remove/item/cart", h.RemoveItemFromCart)
	cart.GET("/my/cart", h.FindCartByUser)
}

func (h *cartHandler) StoreItemToCart(c *gin.Context) {
	var request cart.CartRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Add to Cart", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	authUser := c.MustGet("authUser").(entity.User)

	request.UserID = authUser.ID

	_, err = h.cartService.StoreCart(request)

	if err != nil {
		response := helper.APIResponse("Failed Add to Cart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Add to Successfully.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *cartHandler) RemoveItemFromCart(c *gin.Context) {
	var request cart.CartRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Remove from Cart", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	authUser := c.MustGet("authUser").(entity.User)

	request.UserID = authUser.ID

	_, err = h.cartService.RemoveItemCart(request)

	if err != nil {
		response := helper.APIResponse("Failed Remove from Cart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Remove cart Successfully.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *cartHandler) FindCartByUser(c *gin.Context) {
	var request cart.CartUserIDRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		response := helper.APIResponse("Failed get cart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	authUser := c.MustGet("authUser").(entity.User)

	request.ID = authUser.ID

	getCart, err := h.cartService.FindCartByID(request)
	if err != nil {
		response := helper.APIResponse("Failed get cart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List cart", http.StatusOK, "success", cart.FormatCart(getCart))
	c.JSON(http.StatusOK, response)
}
