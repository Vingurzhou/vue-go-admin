package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	token := c.Query("token")
	fmt.Println(token)
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": struct {
		Roles        []string `json:"roles"`
		Introduction string   `json:"introduction"`
		Avatar       string   `json:"avatar"`
		Name         string   `json:"name"`
	}{Roles: []string{"admin"}, Introduction: "I am a super administrator", Avatar: "https://avatars.githubusercontent.com/u/57127283?v=4", Name: "Vingur Zhou"}})
}
