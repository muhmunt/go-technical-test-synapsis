package routes

import (
	"go-technical-test-synapsis/src/cart"
	"go-technical-test-synapsis/src/handler"
	"go-technical-test-synapsis/src/middleware"
	"go-technical-test-synapsis/src/order"
	"go-technical-test-synapsis/src/product"
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

	productRepository := repository.NewProduct(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProduct(productService)
	productHandler.CustomerMount(routeUser)

	cartRepository := repository.NewCart(db)
	cartItemsRepository := repository.NewCartItems(db)
	cartService := cart.NewService(cartRepository, cartItemsRepository)
	cartHandler := handler.NewCart(cartService)
	cartHandler.Mount(routeUser)

	orderRepository := repository.NewOrder(db)
	orderItemsRepository := repository.NewOrderItems(db)
	orderService := order.NewService(orderRepository, cartItemsRepository, orderItemsRepository)
	orderHandler := handler.NewOrder(orderService)
	orderHandler.Mount(routeUser)
}
