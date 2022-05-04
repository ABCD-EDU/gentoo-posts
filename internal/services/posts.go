package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abcd-edu/gentoo-posts/internal/models"
	"github.com/gin-gonic/gin"
)

func HandlePostContent(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		fmt.Println(err)
		fmt.Printf("ERROR WRITING TO DB: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
	}

	post.CreatedOn = time.Now()
	postId, err := models.WritePost(post)
	if err != nil {
		fmt.Printf("ERROR WRITING TO DB: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	postInfo := models.PostSchema{PostId: postId, PostInfo: post}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"post":   postInfo,
	})
}

func GetUserTimeline(c *gin.Context) {
	userId := c.Query("user_id")

	posts, err := models.GetPostsFromUser(userId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"posts":   posts,
	})
}
