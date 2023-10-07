package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/israel-duff/ledger-system/pkg/controllers"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	authRoutes := r.Group("/auth")

	authRoutes.POST("/register", controllers.RegisterController)
	authRoutes.POST("/login", controllers.LoginController)
}
