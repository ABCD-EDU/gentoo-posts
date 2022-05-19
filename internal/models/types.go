package models

import "time"

type TempUser struct {
	UserId    string    `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	UserInfo  User      `form:"user_info" json:"user_info" xml:"user_info"  binding:"required"`
	CreatedOn time.Time `form:"created_on" json:"created_on" xml:"created_on"  binding:"required"`
	CanPost   bool      `form:"can_post" json:"can_post" xml:"can_post"  binding:"required"`
	IsAdmin   bool      `form:"is_admin" json:"is_admin" xml:"is_admin"`
}

type UserSchema struct {
	UserId   string `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	UserInfo User   `form:"user_info" json:"user_info" xml:"user_info"  binding:"required"`
}

type User struct {
	Username    string `form:"username" json:"username" xml:"username"  binding:"required"`
	Email       string `form:"email" json:"email" xml:"email"  binding:"required"`
	Photo       string `form:"google_photo" json:"google_photo" xml:"google_photo"  binding:"required"`
	Description string `form:"description" json:"description" xml:"description"`
	IsAdmin     bool   `form:"is_admin" json:"is_admin" xml:"is_admin"`
}

type PostResponse struct {
	User       UserSchema `form:"user" json:"user" xml:"user"  binding:"required"`
	Post       PostSchema `form:"post" json:"post" xml:"post"  binding:"required"`
	HateScores Metrics    `form:"hate_scores" json:"hate_scores" xml:"hate_scores"  binding:"required"`
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

type Metrics struct {
	HateScore      float32 `form:"hate_score" json:"hate_score" xml:"hate_score"  binding:"required"`
	NormalScore    float32 `form:"normal_score" json:"normal_score" xml:"normal_score"  binding:"required"`
	OffensiveScore float32 `form:"offensive_score" json:"offensive_score" xml:"offensive_score"  binding:"required"`
	ProfanityScore float32 `form:"profanity_score" json:"profanity_score" xml:"profanity_score"  binding:"required"`
	RaceScore      float32 `form:"race_score" json:"race_score" xml:"race_score"  binding:"required"`
	ReligionScore  float32 `form:"religion_score" json:"religion_score" xml:"religion_score"  binding:"required"`
	SexScore       float32 `form:"sex_score" json:"sex_score" xml:"sex_score"  binding:"required"`
	OtherScore     float32 `form:"other_score" json:"other_score" xml:"other_score"  binding:"required"`
	NoneScore      float32 `form:"none_score" json:"none_score" xml:"none_score"  binding:"required"`
}

type Model struct {
	TaskId int          `form:"task_id" json:"task_id" xml:"task_id"  binding:"required"`
	Scores []ModelScore `form:"scores" json:"scores" xml:"scores"  binding:"required"`
}

type ModelScore struct {
	Class string  `form:"class_value" json:"class_value" xml:"class_value"  binding:"required"`
	Score float32 `form:"score" json:"score" xml:"score"  binding:"required"`
}
