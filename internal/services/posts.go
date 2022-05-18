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
	fmt.Println(c.Query("user_id"))
	if err := c.ShouldBindJSON(&post); err != nil {
		fmt.Printf("ERROR WRITING TO DB: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
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

func GetPost(c *gin.Context) {
	postId := c.Query("post_id")
	userId := c.Query("user_id")

	post, err := models.GetPostByID(userId, postId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"post":    post,
	})
}

func GetUserPosts(c *gin.Context) {
	userId := c.Query("user_id")
	authId := c.Query("auth_id")
	offset := c.Query("offset")
	limit := c.Query("limit")

	postQuery := `
	SELECT 
		u.user_id,
		u.email,
		u.username,
		u.google_photo,
		u.is_admin,

		p.post_id,

		p.content,
		p.created_on,

		m.hate_score,

		m.normal_score,
		m.offensive_score,
		m.profanity_score,
		m.race_score,
		m.religion_score,
		m.sex_score,
		m.other_score,
		m.none_score	
	FROM posts p
	INNER JOIN metrics m
		ON p.post_id=m.post_id
	INNER JOIN users u
		ON u.user_id=p.user_id
	WHERE u.user_id=$1
	ORDER BY p.created_on DESC
	OFFSET $2 LIMIT $3
	`

	posts, err := models.GetPosts(userId, authId, offset, limit, postQuery, "profile")
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

func GetUserTimeline(c *gin.Context) {
	userId := c.Query("user_id")
	authId := c.Query("auth_id")
	offset := c.Query("offset")
	limit := c.Query("limit")

	postQuery := `
	SELECT 
		u.user_id,
		u.email,
		u.username,
		u.google_photo,
		u.is_admin,

		p.post_id,
		p.content,
		p.created_on,

		m.hate_score,

		m.normal_score,
		m.offensive_score,
		m.profanity_score,
		m.race_score,
		m.religion_score,
		m.sex_score,
		m.other_score,
		m.none_score	
	FROM posts p
	INNER JOIN metrics m
		ON p.post_id=m.post_id
	INNER JOIN users u
		ON u.user_id=p.user_id
	WHERE u.user_id IN (
		SELECT followed_id AS user_id
		FROM social_graph
		WHERE follower_id=$1
	) OR u.user_id=$1
	ORDER BY p.created_on DESC
	OFFSET $2 LIMIT $3;
	`

	posts, err := models.GetPosts(userId, authId, offset, limit, postQuery, "timeline")
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

func GetLatestPosts(c *gin.Context) {
	userId := c.Query("user_id")
	authId := c.Query("auth_id")

	postQuery := `
	SELECT 
		u.user_id,
		u.email,
		u.username,
		u.google_photo,
		u.is_admin,

		p.post_id,
		p.content,
		p.created_on,

		m.hate_score,
		m.normal_score,
		m.offensive_score,
		m.profanity_score,
		m.race_score,
		m.religion_score,
		m.sex_score,
		m.other_score,
		m.none_score	
	FROM posts p
	INNER JOIN metrics m
		ON p.post_id=m.post_id
	INNER JOIN users u
		ON u.user_id=p.user_id
	WHERE u.user_id!=$1
	ORDER BY p.created_on DESC
	LIMIT 20;
	`

	posts, err := models.GetPosts(userId, authId, "0", "20", postQuery, "latest")
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
