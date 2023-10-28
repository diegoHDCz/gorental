package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/category", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GET CATEGORY",
			})
		})
		v1.POST("/category", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GET CATEGORY",
			})
		})
		v1.GET("/categories", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GET CATEGORY",
			})
		})
		v1.PUT("/category", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GET CATEGORY",
			})
		})
		v1.DELETE("/category", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GET CATEGORY",
			})
		})
	}
}
