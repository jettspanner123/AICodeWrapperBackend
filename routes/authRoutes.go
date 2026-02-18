package routes

import (
	gin "github.com/gin-gonic/gin"
	controllers "github.com/jettspanner123/AICodeWrapperBackend/controllers"
)

func RegisterAuthRoutes(engine *gin.Engine) {
	userController := controllers.NewAuthController()

	userRoutes := engine.Group("/api/v1/auth")
	{
		userRoutes.GET("/health-check", userController.HealthCheck)
	}
}
