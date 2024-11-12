package router

import (
	"github.com/MarcoTulioFM/controle-gastos/api/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	//ROUTES  ITENS
	itens := router.Group("/v1/api/itens")
	{
		itens.POST("/createItens", handlers.CreateItens)
		itens.GET("/getItens", handlers.GetItens)
		itens.GET("/getItensById",)
		itens.PUT("/updateItens",)
		itens.DELETE("/deleteItens")
	}

	// routes walmart
	walmart := router.Group("/v1/api/walmart")
	{
		walmart.POST("/createwalmart", func(ctx *gin.Context){
			ctx.JSON(200, "msg: OI")
		})
		walmart.GET("/getwalmart",)
		walmart.GET("/getwalmartById",)
		walmart.PUT("/updatewalmart",)
		walmart.DELETE("/deletewalmart")
	}

	// routes itens val
	itensVal := router.Group("/v1/api/itensVal")
	{
		itensVal.POST("/createitensVal", func(ctx *gin.Context){
			ctx.JSON(200, "msg: OI")
		})
		itensVal.GET("/getitensVal",)
		itensVal.GET("/getitensValById",)
		itensVal.PUT("/updateitensVal",)
		itensVal.DELETE("/deleteitensVal")
	}
}