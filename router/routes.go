package router

import (
	"github.com/diegoHDCz/gorental/handler/categories"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/category", categories.ShowCategoryHandler)
		v1.POST("/category", categories.CreateCategoryHandler)
		v1.GET("/categories", categories.ListCategoriesHandler)
		v1.PUT("/category", categories.UpdateCategoryHandler)
		v1.DELETE("/category", categories.DeleteCategoryHandler)
	}
}
