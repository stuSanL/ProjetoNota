package handler

import (
	"ProjetoNota/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotasHandler(ctx *gin.Context) {
	var notas []schemas.Nota
	if err := db.Find(&notas).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprint("Erro ao listar notas"))
		return
	}
	SendSuccess(ctx, "get-notas", notas)
}
