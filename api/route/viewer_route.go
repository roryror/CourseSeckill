package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)	

func NewViewerRoute(router *gin.Engine) {
	viewer := router.Group("/viewer")
	{
		viewer.GET("/orders", handleGetAllOrders)
		viewer.GET("/courses", handleGetAllCourses)
		viewer.GET("/stock", handleGetAllStock)
		viewer.GET("/order-status", handleGetAllOrderStatus)
	}
}

func handleGetAllOrders(c *gin.Context) {
	orders, err := controller.ViewerController.GetAllOrders(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get all orders", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "orders": orders})
}

func handleGetAllCourses(c *gin.Context) {
	courses, err := controller.ViewerController.GetAllCourses(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get all courses", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "courses": courses})
}

func handleGetAllStock(c *gin.Context) {
	stock, err := controller.ViewerController.GetAllStock(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get all stock", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "stock": stock})
}

func handleGetAllOrderStatus(c *gin.Context) {
	orderStatus, err := controller.ViewerController.GetAllOrderStatus(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get all order status", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "orderStatus": orderStatus})
}

