package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": struct {
		Token string `json:"token"`
	}{Token: "admin-token"}})
}
