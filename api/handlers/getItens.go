package handlers

import (
	"net/http"

	"github.com/MarcoTulioFM/controle-gastos/api/config"
	"github.com/MarcoTulioFM/controle-gastos/api/schemas"
	"github.com/gin-gonic/gin"
)

// getItens busca todos os itens no banco de dados e retorna em formato JSON
func GetItens(ctx *gin.Context) {
	var itens []schemas.Itens

	// Busca todos os itens no banco
	if err := config.DB.Find(&itens).Error; err != nil {
		// Retorna erro se falhar ao buscar os itens
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar itens",
		})
		return
	}

	// Retorna os itens encontrados
	ctx.JSON(http.StatusOK, gin.H{
		"itens": itens,
	})
}
