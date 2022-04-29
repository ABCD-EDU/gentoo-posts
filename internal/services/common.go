package services

import (
	"fmt"
	"net/http"

	"github.com/abcd-edu/gentoo-posts/internal/models"
	"github.com/gin-gonic/gin"
)

func HandleMain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"api": "Post Services",
	})
}

func HandleRegistration(c *gin.Context) {
	var user models.UserSchema
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("ERROR WITH PARSING JSON BODY")
	}

	fmt.Println(user)

	_, err := models.WriteUserRegistration(user)
	if err != nil {
		fmt.Println("ERROR WRITING TO DB")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func HandlePost(c *gin.Context) {
	//	userId := c.Query("user_id")
	//	content := c.Query("content")
	//	createdOn := time.Now()
	//
	//	var post models.Post
	//	if err := c.ShouldBindJSON(&post); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//	}
	//
	//	models.WritePost()
}
