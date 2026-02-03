package handler

import (
	"ProjetoNota/schemas"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Adiciona nota
// @Description Adiciona uma nova nota
// @Tags Notas
// @Accept json
// @Produce json
// @Param request body AddNotaRequest true "Request Body"
// @Success 200 {object} AddNotaResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /nota/add [post]
func AddNotaHandler(ctx *gin.Context) {
	request := AddNotaRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		fmt.Printf("Erro de validação: %s", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	nota := schemas.Nota{
		Data:      time.Now(),
		Pontuacao: *request.Pontuacao,
		Texto:     request.Texto,
	}

	if err := db.Create(&nota).Error; err != nil {
		fmt.Printf("Erro criando nota: %s", err.Error())
		return
	}

	SendSuccess(ctx, "add-nota", nota)

}
