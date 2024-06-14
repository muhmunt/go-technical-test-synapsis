package routes

import (
	"go-technical-test-synapsis/src/handler"
	"go-technical-test-synapsis/src/middleware"
	"go-technical-test-synapsis/src/repository"
	"go-technical-test-synapsis/src/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	routeAuth := r.Group("/auth")

	routeUser := r.Group("/user")

	userRepository := repository.NewUser(db)
	userService := user.NewService(userRepository)

	// authorization
	routeUser.Use(middleware.Authorization(userService, "USER"))

	userHandler := handler.NewUser(userService)
	userHandler.Mount(routeAuth)
}
