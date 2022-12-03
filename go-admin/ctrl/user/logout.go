package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": "success"})
}
