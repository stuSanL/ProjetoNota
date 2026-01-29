package router

import (
	docs "ProjetoNota/docs"
	"ProjetoNota/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {

	handler.Initialize()

	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath

	v := router.Group(basePath)
	{
		v.GET("/notas", handler.GetNotasHandler)
		v.GET("/nota", handler.GetNotaHandler)

		v.POST("/nota/add", handler.AddNotaHandler)

		v.DELETE("/nota/delete", handler.DeleteNotaHandler)

		v.PUT("/nota/update", handler.UpdateNotaHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
