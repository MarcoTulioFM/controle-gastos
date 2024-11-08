package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteITem(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Delete Succeful",
	})
}