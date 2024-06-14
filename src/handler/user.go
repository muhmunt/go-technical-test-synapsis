package handler

import (
	"go-technical-test-synapsis/src/auth"
	"go-technical-test-synapsis/src/helper"
	"go-technical-test-synapsis/src/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUser(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Mount(user *gin.RouterGroup) {
	user.POST("/register", h.RegisterUser)
	user.POST("/login", h.LoginUser)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var request user.RegisterRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed register account.", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvailableUsername, err := h.userService.CheckUsernameAvailable(request.Username)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Username checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if !isAvailableUsername {
		errorMessage := gin.H{"errors": "Username is already taken"}
		response := helper.APIResponse("Failed register account.", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(request)

	if err != nil {
		response := helper.APIResponse("Failed register account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := auth.GenerateToken(newUser.ID, newUser.Role)

	if err != nil {
		response := helper.APIResponse("Failed register account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account registered successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var request user.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.ValidationFormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userData, err := h.userService.Login(request)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := auth.GenerateToken(userData.ID, userData.Role)

	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(userData, token)

	response := helper.APIResponse("Login successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
