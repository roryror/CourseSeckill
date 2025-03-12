package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)	

func NewWarmupRoute(router *gin.Engine) {
	warmup := router.Group("/warmup")
	{
		warmup.POST("/", handleWarmup)
	}
}

func handleWarmup(c *gin.Context) {
	err := controller.InitController.WarmupCache(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "warmup failed", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "warmup success"})
}