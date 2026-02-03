package handler

import (
	"ProjetoNota/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotasHandler @Summary Obter notas
// @Description Lista todas as notas
// @Tags Notas
// @Accept json
// @Produce json
// @Success 200 {array} schemas.NotaResponse
// @Failure 500 {object} ErrorResponse
// @Router /notas [get]
func GetNotasHandler(ctx *gin.Context) {
	var notas []schemas.Nota
	if err := db.Find(&notas).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprint("Erro ao listar notas"))
		return
	}
	SendSuccess(ctx, "get-notas", notas)
}
