package server

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	router := gin.Default()

	port := os.Getenv("PORT")

	api := router.Group("/api")

	fmt.Print(api)
	router.Run(":" + port)
}
