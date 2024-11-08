package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update Succeful",
	})
}