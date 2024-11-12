package handlers

import (
	"net/http"

	"github.com/MarcoTulioFM/controle-gastos/api/config"
	"github.com/MarcoTulioFM/controle-gastos/api/schemas"
	"github.com/gin-gonic/gin"
)

// createItens cria um novo item no banco de dados
func CreateItens(ctx *gin.Context) {
	var newItem schemas.Itens

	// Vincula os dados JSON à estrutura Itens
	if err := ctx.ShouldBindJSON(&newItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Salva o novo item no banco de dados
	if err := config.DB.Create(&newItem).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao salvar o item"})
		return
	}

	// Responde com sucesso
	ctx.JSON(http.StatusCreated, gin.H{"message": "Item criado com sucesso", "data": newItem})
}
