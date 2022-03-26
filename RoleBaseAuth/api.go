package main

import (
	handler "RoleBaseAuth/handlers"
	"RoleBaseAuth/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/login", handler.LoginHandler)

	api := route.Group("/api")
	api.Use(middleware.ValidateToken())

	product := api.Group("/product")

	product.Use(middleware.Authorization([]int{1}))
	product.GET("/", handler.GetAll)
	product.POST("/", middleware.Authorization([]int{4}), handler.AddProduct)

	user := api.Group("/User")
	user.GET("/", func(c *gin.Context) {
		c.AbortWithStatusJSON(200, gin.H{
			"status": "ok",
		})
	})
	route.Run("localhost:8080")
}
