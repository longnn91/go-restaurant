package main

import (
	"gogo/modules/auth/middleware"
	auth "gogo/modules/auth/transport"

	// handlers "gogo/internal/http/server/handlers"

	category "gogo/modules/category/controller"
	food "gogo/modules/food/controller"
	menu "gogo/modules/menu/controller"
	order "gogo/modules/order/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter initializes and returns a Gin Engine with all the routes configured
func SetupRouter(db *gorm.DB) *gin.Engine {
	// userRepo := &usecases.UserUsecase{Repo: repository.UserRepository{DB: db}}

	r := gin.Default()

	v1 := r.Group("/v1")

	foodRouter := v1.Group("/food").Use(middleware.AuthMiddleware)
	{
		foodRouter.POST("/", food.CreateFood(db))
		foodRouter.GET("/", food.GetFoods(db))
		foodRouter.GET("/:id", food.GetFoodById(db))
		foodRouter.PUT("/:id", food.UpdateFood(db))
		foodRouter.DELETE("/:id", food.DeleteFood(db))
	}

	menuRouter := v1.Group("/menu").Use(middleware.AuthMiddleware)
	{
		menuRouter.POST("/", menu.CreateMenu(db))
		menuRouter.GET("/", menu.GetMenus(db))
		menuRouter.GET("/:id", menu.GetMenuById(db))
		menuRouter.PUT("/:id", menu.UpdateMenu(db))
		menuRouter.DELETE("/:id", menu.DeleteMenu(db))
	}

	categoryRouter := v1.Group("/category").Use(middleware.AuthMiddleware)
	{
		categoryRouter.POST("/", category.CreateCategory(db))
		categoryRouter.GET("/", category.GetCategories(db))
		categoryRouter.GET("/:id", category.GetCategoryById(db))
		categoryRouter.PUT("/:id", category.UpdateCategory(db))
		categoryRouter.DELETE("/:id", category.DeleteCategory(db))
	}

	orderRouter := v1.Group("/order").Use(middleware.AuthMiddleware)
	{
		orderRouter.POST("/", order.CreateOrder(db))
		orderRouter.GET("/", order.GetOrders(db))
		orderRouter.GET("/:id", order.GetOrderById(db))
		orderRouter.PUT("/:id", order.UpdateOrder(db))
		orderRouter.DELETE("/:id", order.DeleteOrder(db))
	}

	authRouter := v1.Group("/auth")
	{
		authRouter.POST("/login", auth.Login(db))
		authRouter.POST("/register", auth.Register(db))
		// authRouter.POST("/login", handlers.Login(userRepo))
		// authRouter.POST("/register", handlers.Register(userRepo))
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
