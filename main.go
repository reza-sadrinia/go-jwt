package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	routes "github.com/reza-sadrinia/go_jwt/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()

	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "access granted to api 1",
		})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "access granted to api 2",
		})
	})

	router.Run(":" + port)
}
