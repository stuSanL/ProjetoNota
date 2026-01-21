package router

import (
	"ProjetoNota/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	handler.Initialize()

	v := router.Group("/api/")
	{
		v.GET("/notas", handler.GetNotasHandler)
		v.GET("/nota/:id", handler.GetNotaHandler)
		v.POST("/addNota", handler.AddNotaHandler)
	}

}
