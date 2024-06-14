package routes

import (
	"go-technical-test-synapsis/src/middleware"
	"go-technical-test-synapsis/src/repository"
	"go-technical-test-synapsis/src/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAdminRoutes(r *gin.RouterGroup, db *gorm.DB) {

	routeAdmin := r.Group("/admin")

	userRepository := repository.NewUser(db)
	userService := user.NewService(userRepository)

	// authorization
	routeAdmin.Use(middleware.Authorization(userService, "ADMIN"))

}
