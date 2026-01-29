package handler

import (
	"ProjetoNota/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteNotaHandler(ctx *gin.Context) {
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

	if err := db.Delete(&nota).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Erro ao deletar nota %s", id))
		return
	}

	SendSuccess(ctx, "delete-nota", id)
}
