package routes

import (
	"ewallet-blockhain/app/config"

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
	// productHandlerV1 := productHandlerV1.Handler(db)
	// productViewV1 := productviewV1.View(db)
	//blockhain := blockhain.Init(conf)

	// Routing Website Service
	//product := router.Group("/product", basic.Auth(conf))
	//product.GET("/", productViewV1.Index)

	//Routing API Service
	api := router.Group("/api/v1")
	//api.GET("/product", productHandlerV1.ListProduct)

	router = ParseTmpl(router)
	return router
}
