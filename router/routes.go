package router

import (
	"github.com/diegoHDCz/gorental/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/category", handler.ShowCategoryHandler)
		v1.POST("/category", handler.CreateCategoryHandler)
		v1.GET("/categories", handler.ListCategoriesHandler)
		v1.PUT("/category", handler.UpdateCategoryHandler)
		v1.DELETE("/category", handler.DeleteCategoryHandler)
	}
}
