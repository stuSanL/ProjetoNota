package controller

import (
	"ProjetoNota/model"
	"ProjetoNota/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type notaController struct {
	NotaService service.NotaService
}

func NewNotaController(service service.NotaService) notaController {
	return notaController{
		NotaService: service,
	}
}

func (n *notaController) GetNotas(c *gin.Context) {
	notas, err := n.NotaService.GetNotas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, notas)
}

func (n *notaController) AddNota(c *gin.Context) {
	var nota model.Nota
	err := c.BindJSON(&nota)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	insertedNota, err := n.NotaService.AddNota(nota)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, insertedNota)
}

func (n *notaController) GetNota(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "id is empty")
		return
	}

	notaId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	nota, err := n.NotaService.GetNota(notaId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if nota == nil {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	c.JSON(http.StatusOK, nota)
}
