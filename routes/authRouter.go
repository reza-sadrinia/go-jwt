package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/reza-sadrinia/go_jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
