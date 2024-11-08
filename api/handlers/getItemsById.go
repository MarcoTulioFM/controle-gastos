package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItemsById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "List by item Succeful",
	})
}