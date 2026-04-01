package router

import (
	"thoropa/internal/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(linkHandler *handler.LinkHandler) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0/0"})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API rodando 🚀"})
	})

	r.POST("/link", linkHandler.Create)
	r.GET("/link/:id", linkHandler.GetById)
	r.DELETE("/link/:id", linkHandler.DeleteById)
	r.GET("/links", linkHandler.GetByIP)

	return r
}
