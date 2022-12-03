package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"zwzProject/ctrl/article"
	"zwzProject/ctrl/user"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}
func main() {
	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	//解决跨域
	router.Use(Cors())
	// Simple group: v1
	v1 := router.Group("/article")
	{
		v1.GET("/list", article.List)
		v1.GET("/detail", article.List)
		v1.GET("/pv", article.List)
		v1.POST("/create", article.List)
		v1.POST("/update", article.List)
	}

	// Simple group: v2
	v2 := router.Group("/user")
	{
		v2.POST("/login", user.Login)
		v2.GET("/info", user.Info)
		v2.POST("/logout", user.Logout)
	}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
