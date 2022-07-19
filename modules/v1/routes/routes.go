package routes

import (
	"ewallet-blockhain/app/blockhain"
	"ewallet-blockhain/app/config"
	walletHandlerV1 "ewallet-blockhain/modules/v1/utilities/wallet/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ParseTmpl(router *gin.Engine) *gin.Engine { //Load HTML Template
	router.Static("/assets", "./public/assets")
	router.Static("/images", "./public/images")
	router.Static("/css", "./public/assets/css")
	router.Static("/js", "./public/assets/js")
	router.Static("/fonts", "./public/assets/fonts")
	return router
}

func Init(db *gorm.DB, conf config.Conf, router *gin.Engine) *gin.Engine {
	blockhain := blockhain.Init(conf)
	walletHandlerV1 := walletHandlerV1.Handler(db, blockhain)

	//Routing API Service
	api := router.Group("/api/v1")
	api.GET("/balance", walletHandlerV1.GetBalance)
	api.GET("/mywallet", walletHandlerV1.GetMyWallet)

	router = ParseTmpl(router)
	return router
}
