package main

import (
	"github.com/abcd-edu/gentoo-posts/internal/configs"
	"github.com/abcd-edu/gentoo-posts/internal/models"
	"github.com/abcd-edu/gentoo-posts/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configs.InitializeViper()
	models.InitializeDB()

	router := gin.Default()
	router.Use(CORSMiddleware())

	v1 := router.Group("/v1")
	v1.Use(CORSMiddleware())
	{
		v1.GET("/", services.HandleMain)
		v1.POST("/register", services.HandleRegistration)
		v1.POST("/submit-post", services.HandlePostContent)
		v1.GET("/user/profile", services.GetUserPosts)
		v1.GET("/user/timeline", services.GetUserTimeline)
		v1.POST("/user/mute", services.MuteUser)
		v1.POST("/user/ban", services.BanUser)
		v1.GET("/post", services.GetPost)
		v1.GET("/post/latest", services.GetLatestPosts)
	}

	port := viper.GetString("port")
	router.Run(":" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
