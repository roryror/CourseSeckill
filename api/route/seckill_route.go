package route

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewSeckillRoute(router *gin.Engine) {
	seckill := router.Group("/seckill")
	{
		seckill.POST("/:cid/:uid", handleSeckill)
	}
}

func handleSeckill(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	uid, _ := strconv.Atoi(c.Param("uid"))

	var test bool = true
	if test {
		if cid == 0 {
			cid = rand.Intn(6) + 1
		}
		if uid == 0 {
			uid = rand.Intn(1000) + 1
		}
	}

	err := controller.SeckillController.RunSeckill(ctx, cid, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "seckill failed", "error": err.Error(), "courseID": cid})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "seckill order pending", "error": nil, "courseID": cid})
}
