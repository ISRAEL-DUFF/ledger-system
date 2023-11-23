package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/israel-duff/ledger-system/pkg/controllers"
)

func RegisterWalletRoutes(r *gin.RouterGroup) {
	walletRoutes := r.Group("/wallet")

	walletController := controllers.NewWalletController()

	walletRoutes.POST("/create", walletController.CreateWallet)
	walletRoutes.GET("/get/:accountNumber", walletController.GetWalletByAccountNumber)
	walletRoutes.GET("/status/:accountNumber", walletController.GetWalletStatus)
	walletRoutes.GET("/list/:ownerId", walletController.ListUserWalllets)
	walletRoutes.GET("/ledger-account/:accountNumber", walletController.GetLedgerAccount)
	walletRoutes.POST("/type/create", walletController.CreateWalletType)
	walletRoutes.POST("/type/update", walletController.UpdateWalletType)
	walletRoutes.GET("/type/list/:ownerId", walletController.ListWalletTypes)
	walletRoutes.POST("/transaction/post", walletController.PostTransaction)
	walletRoutes.GET("/account-statement/:accountNumber", walletController.GetAccountStatement)

}
