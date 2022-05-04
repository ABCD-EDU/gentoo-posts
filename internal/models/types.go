package models

import "time"

type UserSchema struct {
	UserId   string `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	UserInfo User   `form:"user_info" json:"user_info" xml:"user_info"  binding:"required"`
}

type User struct {
	Username string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Photo    string `form:"google_photo" json:"google_photo" xml:"google_photo"  binding:"required"`
}

type PostSchema struct {
	PostId   string `form:"post_id" json:"post_id" xml:"post_id"  binding:"required"`
	PostInfo Post   `form:"post_info" json:"post_info" xml:"post_info"  binding:"required"`
}

type Post struct {
	UserId    string    `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Content   string    `form:"content" json:"content" xml:"content"  binding:"required"`
	CreatedOn time.Time `form:"created_on" json:"created_on" xml:"created_on"  binding:"required"`
}

type Like struct {
	PostId int `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Amount int `form:"amount" json:"amount" xml:"amount"  binding:"required"`
}

type Model struct {
	TaskId int          `form:"task_id" json:"task_id" xml:"task_id"  binding:"required"`
	Scores []ModelScore `form:"scores" json:"scores" xml:"scores"  binding:"required"`
}

type ModelScore struct {
	Class string  `form:"class_value" json:"class_value" xml:"class_value"  binding:"required"`
	Score float32 `form:"score" json:"score" xml:"score"  binding:"required"`
}
