package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": "success"})
}
