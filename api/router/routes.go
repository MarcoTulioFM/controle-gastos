package router

import (
	"github.com/MarcoTulioFM/controle-gastos/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1/items",)
	{
		v1.GET("/getItems", handlers.GetItems)
		v1.GET("/getItemsById", handlers.GetItemsById)
		v1.POST("/createItem", handlers.CreateItem)
		v1.DELETE("/deleteItem", handlers.DeleteITem)
		v1.PUT("/updateItem", handlers.UpdateItems)

	} 
}