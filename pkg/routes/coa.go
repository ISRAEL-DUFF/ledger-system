package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/israel-duff/ledger-system/pkg/controllers"
)

func RegisterCOARoutes(r *gin.RouterGroup) {
	authRoutes := r.Group("/coa")

	authRoutes.GET("/", controllers.ListAccounts)
	authRoutes.POST("/account", controllers.CreateAccount)
}
