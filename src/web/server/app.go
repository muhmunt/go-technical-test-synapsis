package server

import (
	"go-technical-test-synapsis/src/web/routes"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	router := gin.Default()

	port := os.Getenv("PORT")

	api := router.Group("/api")

	routes.SetupUserRoutes(api, db)
	routes.SetupAdminRoutes(api, db)

	router.Run(":" + port)
}
