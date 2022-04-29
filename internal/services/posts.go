package services

import (
	"fmt"
	"net/http"

	"github.com/abcd-edu/gentoo-posts/internal/models"
	"github.com/gin-gonic/gin"
)

func HandlePostContent(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		fmt.Println(err)
	}

	postId, err := models.WritePost(post)
	if err != nil {
		fmt.Printf("ERROR WRITING TO DB: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
	}

	postInfo := models.PostSchema{PostId: postId, PostInfo: post}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"post":   postInfo,
	})
}
