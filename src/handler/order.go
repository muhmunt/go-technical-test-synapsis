package handler

import (
	"go-technical-test-synapsis/src/entity"
	"go-technical-test-synapsis/src/helper"
	"go-technical-test-synapsis/src/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	orderService order.Service
}

func NewOrder(orderService order.Service) *orderHandler {
	return &orderHandler{orderService}
}

func (h *orderHandler) Mount(order *gin.RouterGroup) {
	order.POST("/checkout", h.CreateOrder)
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var request order.OrderRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	authUser := c.MustGet("authUser").(entity.User)
	request.UserID = authUser.ID

	_, err = h.orderService.CreateOrder(request)

	if err != nil {
		response := helper.APIResponse("Failed create order "+err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	response := helper.APIResponse("Order created successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
