package route

import (
	"github.com/gin-gonic/gin"
	"lecture/1230/controller"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/mnemonics", controller.NewMnemonic)
	router.POST("/wallets", controller.NewWallet)
	router.POST("/wallets/keystores", controller.NewWalletWithKeystore)
	return router
}