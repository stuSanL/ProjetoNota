package handler

import (
	"ProjetoNota/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNotaHandler @Summary Obter nota
// @Description Acha uma nota conforme id
// @Tags Notas
// @Accept json
// @Produce json
// @Param id query int true "Id da nota a ser buscada"
// @Success 200 {object} schemas.NotaResponse
// @Failure 400 {object} ErrorResponse
// @Router /nota [get]
func GetNotaHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errRequiredParam("id", "queryParam").Error())
		return
	}

	nota := schemas.Nota{}

	if err := db.First(&nota, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("nota %s not found", id))
		return
	}

	SendSuccess(ctx, "get-nota-"+id, nota)
}
