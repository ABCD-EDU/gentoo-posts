package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetPosts(userId string, authId string, offset string, limit string, hateFilter string, postQuery string, queryType string) ([]PostResponse, error) {
	posts := []PostResponse{}

	userQuery := `
	SELECT is_admin FROM users
	WHERE user_id=$1
	LIMIT 1`

	var isAdmin bool
	if err := db.QueryRow(userQuery, authId).Scan(&isAdmin); err != nil {
		return posts, err
	}

	var rows *sql.Rows
	var err error
	if queryType == "timeline" {
		rows, err = db.Query(postQuery, userId, hateFilter, offset, limit)
	} else if queryType == "profile" {
		rows, err = db.Query(postQuery, userId, hateFilter, offset, limit)
	} else if queryType == "latest" {
		rows, err = db.Query(postQuery, userId, hateFilter)
	}
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		// USER VARS
		var userId, email, username, googlePhoto string
		var userAdmin bool
		// POST VARS
		var postId, content string
		var createdOn time.Time
		// METRICS SCORE
		var hateScore, normalScore, offnScore, prfnScore, raceScore, religionScore, sexScore, otherScore, noneScore float32
		err := rows.Scan(
			&userId,
			&email,
			&username,
			&googlePhoto,
			&userAdmin,
			&postId,
			&content,
			&createdOn,
			&hateScore,
			&normalScore,
			&offnScore,
			&prfnScore,
			&raceScore,
			&religionScore,
			&sexScore,
			&otherScore,
			&noneScore,
		)
		if err != nil {
			return posts, err
		}

		// USER STRUCTS
		userInfo := &User{Username: username, Email: email, Photo: googlePhoto, IsAdmin: userAdmin}
		user := &UserSchema{UserId: userId, UserInfo: *userInfo}

		// METRIC STRUCTS
		metrics := &Metrics{
			HateScore:      hateScore,
			NormalScore:    normalScore,
			OffensiveScore: offnScore,
			ProfanityScore: prfnScore,
			RaceScore:      raceScore,
			ReligionScore:  religionScore,
			SexScore:       sexScore,
			OtherScore:     otherScore,
			NoneScore:      noneScore,
		}

		// POST STRUCT
		postInfo := &Post{UserId: userId, Content: content, CreatedOn: createdOn}
		post := &PostSchema{PostId: postId, PostInfo: *postInfo}

		var postResponse PostResponse
		if isAdmin {
			postResponse = PostResponse{User: *user, Post: *post, HateScores: *metrics}
		} else {
			postResponse = PostResponse{User: *user, Post: *post, HateScores: Metrics{}}
		}
		posts = append(posts, postResponse)
	}

	return posts, nil

}

func GetPostByID(authId string, postId string) (PostResponse, error) {
	post := PostResponse{}

	userQuery := `
	SELECT is_admin FROM users
	WHERE user_id=$1
	LIMIT 1`

	var isAdmin bool
	if err := db.QueryRow(userQuery, authId).Scan(&isAdmin); err != nil {
		return post, err
	}

	postQuery := `
	SELECT
		u.user_id,
		u.email,
		u.username,
		u.google_photo,
		u.is_admin,

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
	WHERE p.post_id=$1;
	`
	// USER VARS
	var userId, email, username, googlePhoto string
	var userAdmin bool
	// POST VARS
	var content string
	var createdOn time.Time
	// METRICS SCORE
	var hateScore, normalScore, offnScore, prfnScore, raceScore, religionScore, sexScore, otherScore, noneScore float32

	if err := db.QueryRow(postQuery, postId).Scan(
		&userId,
		&email,
		&username,
		&googlePhoto,
		&userAdmin,
		&content,
		&createdOn,
		&hateScore,
		&normalScore,
		&offnScore,
		&prfnScore,
		&raceScore,
		&religionScore,
		&sexScore,
		&otherScore,
		&noneScore,
	); err != nil {
		return post, err
	}

	// USER STRUCTS
	userInfo := &User{Username: username, Email: email, Photo: googlePhoto, IsAdmin: userAdmin}
	user := &UserSchema{UserId: userId, UserInfo: *userInfo}

	// METRIC STRUCTS
	metrics := &Metrics{
		HateScore:      hateScore,
		NormalScore:    normalScore,
		OffensiveScore: offnScore,
		ProfanityScore: prfnScore,
		RaceScore:      raceScore,
		ReligionScore:  religionScore,
		SexScore:       sexScore,
		OtherScore:     otherScore,
		NoneScore:      noneScore,
	}

	// POST STRUCT
	postInfo := &Post{UserId: userId, Content: content, CreatedOn: createdOn}
	postSchema := &PostSchema{PostId: postId, PostInfo: *postInfo}

	if isAdmin {
		post = PostResponse{User: *user, Post: *postSchema, HateScores: *metrics}
	} else {
		post = PostResponse{User: *user, Post: *postSchema, HateScores: Metrics{}}
	}

	return post, nil
}

/*
	TRANSACTION PROCESS:
	1. Get user info from given user_id
	2. Get specific post from given user_id
	3. Get metrics of the post if user_id is admin, else, empty array
*/
func GetPostsFromUser(userId string) ([]PostResponse, error) {
	posts := []PostResponse{}

	userQuery := `
	SELECT is_admin FROM users
	WHERE user_id=$1
	LIMIT 1`
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
	ORDER BY p.created_on DESC;
	`

	var isAdmin bool
	if err := db.QueryRow(userQuery, userId).Scan(&isAdmin); err != nil {
		return posts, err
	}

	rows, err := db.Query(postQuery, userId)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		// USER VARS
		var userId, email, username, googlePhoto string
		var isAdmin bool
		// POST VARS
		var postId, content string
		var createdOn time.Time
		// METRICS SCORE
		var hateScore, normalScore, offnScore, prfnScore, raceScore, religionScore, sexScore, otherScore, noneScore float32

		err := rows.Scan(
			&userId,
			&email,
			&username,
			&googlePhoto,
			&isAdmin,
			&postId,
			&content,
			&createdOn,
			&hateScore,
			&normalScore,
			&offnScore,
			&prfnScore,
			&raceScore,
			&religionScore,
			&sexScore,
			&otherScore,
			&noneScore,
		)
		if err != nil {
			return posts, err
		}

		// USER STRUCTS
		userInfo := &User{Username: username, Email: email, Photo: googlePhoto, IsAdmin: isAdmin}
		user := &UserSchema{UserId: userId, UserInfo: *userInfo}

		// METRIC STRUCTS
		metrics := &Metrics{
			HateScore:      hateScore,
			NormalScore:    normalScore,
			OffensiveScore: offnScore,
			ProfanityScore: prfnScore,
			RaceScore:      raceScore,
			ReligionScore:  religionScore,
			SexScore:       sexScore,
			OtherScore:     otherScore,
			NoneScore:      noneScore,
		}

		// POST STRUCT
		postInfo := &Post{UserId: userId, Content: content, CreatedOn: createdOn}
		post := &PostSchema{PostId: postId, PostInfo: *postInfo}

		var postResponse PostResponse
		if isAdmin {
			postResponse = PostResponse{User: *user, Post: *post, HateScores: *metrics}
		} else {
			postResponse = PostResponse{User: *user, Post: *post, HateScores: Metrics{}}
		}
		posts = append(posts, postResponse)
	}

	return posts, nil
}

func GetHateSpeechScores(content string) ([]Model, error) {
	req, err := http.NewRequest("GET", "http://localhost:8010/models/bert/", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("message", content)
	req.URL.RawQuery = q.Encode()

	res, err := http.Get(req.URL.String())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resBytes := []byte(string(string(body)))
	var jsonRes []Model
	_ = json.Unmarshal(resBytes, &jsonRes)

	return jsonRes, nil
}

func WritePost(post Post) (string, error) {
	sqlQuery := `
	INSERT INTO posts(user_id, content, created_on)
	VALUES ($1, $2, $3)
	RETURNING post_id;
	`
	metricQuery := `
	INSERT INTO metrics(post_id, hate_score, normal_score, offensive_score, profanity_score, race_score, religion_score, sex_score, other_score, none_score)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`

	hateScores, err := GetHateSpeechScores(post.Content)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var postId string
	if err := db.QueryRow(
		sqlQuery,
		post.UserId,
		post.Content,
		post.CreatedOn,
	).Scan(&postId); err != nil {
		fmt.Println("ERROR WRITING TO DB")
		fmt.Println(err)
		return "", err
	}

	_, err = db.Exec(
		metricQuery,
		postId,
		hateScores[0].Scores[0].Score,
		hateScores[0].Scores[1].Score,
		hateScores[0].Scores[2].Score,
		hateScores[1].Scores[0].Score,
		hateScores[2].Scores[0].Score,
		hateScores[2].Scores[1].Score,
		hateScores[2].Scores[2].Score,
		hateScores[2].Scores[3].Score,
		hateScores[2].Scores[4].Score,
	)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	return postId, nil
}

func WriteReplyPost() {

}
