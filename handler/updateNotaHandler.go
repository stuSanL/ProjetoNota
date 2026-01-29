package handler

import (
	"ProjetoNota/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateNotaHandler(ctx *gin.Context) {
	// nota/update?id=n, data -> body
	request := UpdateNotaRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		fmt.Printf("Erro de validação: %s", err)
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errRequiredParam("id", "querryParam").Error())
		return
	}

	nota := schemas.Nota{}
	if err := db.First(&nota, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "nota não achada")
		return
	}

	if request.Pontuacao != nil {
		nota.Pontuacao = *request.Pontuacao
	}
	if request.Texto != "" {
		nota.Texto = request.Texto
	}

	if err := db.Save(&nota).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "Erro atualizando nota "+id+": "+err.Error())
		return
	}
	SendSuccess(ctx, "update-nota-"+id, nota)
}
