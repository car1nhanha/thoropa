package router

import (
	"thoropa/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(linkHandler *handler.LinkHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API rodando 🚀"})
	})

	r.POST("/link", linkHandler.Create)

	return r
}
