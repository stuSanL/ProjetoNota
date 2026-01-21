package handler

import (
	"ProjetoNota/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddNotaHandler(ctx *gin.Context) {
	request := CreateNotaRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		fmt.Printf("Erro de validação: %s", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	nota := schemas.Nota{
		Data:      request.Data,
		Pontuacao: *request.Pontuacao,
		Texto:     request.Texto,
	}

	if err := db.Create(&nota).Error; err != nil {
		fmt.Printf("Erro criando nota: %s", err.Error())
		return
	}

	SendSuccess(ctx, "add-nota", nota)

}
