package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategoriesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}
