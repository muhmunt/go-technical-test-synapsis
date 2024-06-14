package handler

import (
	"go-technical-test-synapsis/src/helper"
	"go-technical-test-synapsis/src/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProduct(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) Mount(product *gin.RouterGroup) {
	product.POST("/product", h.StoreProduct)
	product.GET("/product/:id", h.FindProductByID)
	product.PUT("/product/:id", h.UpdateProduct)
	product.DELETE("/product/:id", h.DeleteProductByID)
	product.GET("/products", h.FindProducts)
}

func (h *productHandler) CustomerMount(customerProduct *gin.RouterGroup) {
	customerProduct.GET("/products", h.FindProducts)
	customerProduct.GET("/products/:category", h.FindProductByCategory)
}

func (h *productHandler) StoreProduct(c *gin.Context) {
	var request product.ProductRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to Store Product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := h.productService.StoreProduct(request)

	if err != nil {
		response := helper.APIResponse("Failed to Store Product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Store Product Successfully.", http.StatusOK, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FindProductByCategory(c *gin.Context) {
	var request product.ProductCategoryRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		response := helper.APIResponse("Failed get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.productService.FindProductByCategory(request)
	if err != nil {
		response := helper.APIResponse("Failed get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FindProducts(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	offset := (page - 1) * pageSize

	products, err := h.productService.FindProducts(offset, pageSize)
	if err != nil {
		response := helper.APIResponse("Failed get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FindProductByID(c *gin.Context) {
	var request product.ProductIDRequest

	err := c.ShouldBindUri(&request)
	if err != nil {
		response := helper.APIResponse("Failed get product by ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getProduct, err := h.productService.FindProductByID(request)

	if err != nil {
		response := helper.APIResponse("Failed get product by ID.", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product Detail", http.StatusOK, "success", product.FormatProduct(getProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var requestID product.ProductIDRequest

	err := c.ShouldBindUri(&requestID)

	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var request product.ProductRequest

	err = c.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateProduct, err := h.productService.UpdateProduct(requestID, request)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Update product successfully.", http.StatusOK, "success", product.FormatProduct(updateProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProductByID(c *gin.Context) {
	var requestID product.ProductIDRequest

	err := c.ShouldBindUri(&requestID)

	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.productService.DeleteProductByID(requestID)

	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete product successfully.", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
