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
	var user models.TempUser
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		fmt.Println("ERROR WITH PARSING JSON BODY")
		return
	}

	fmt.Println(user)

	_, err := models.WriteUserRegistration(user)
	if err != nil {
		fmt.Println("ERROR WRITING TO DB")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

type UserParam struct {
	UserId string `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
}

func MuteUser(c *gin.Context) {
	var param UserParam
	if err := c.ShouldBindJSON(&param); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	err := models.MuteUser(param.UserId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func BanUser(c *gin.Context) {
	var param UserParam
	if err := c.ShouldBindJSON(&param); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	err := models.BanUser(param.UserId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
