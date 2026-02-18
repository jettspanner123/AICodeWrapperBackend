package auth

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"
	responseModels "github.com/jettspanner123/AICodeWrapperBackend/models/response"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (authController *AuthController) HealthCheck(_context *gin.Context) {
	_context.JSON(http.StatusOK, responseModels.APIBaseResponse{
		Success: true,
		Message: "Auth Service Working Fine! ✅",
	})
}
