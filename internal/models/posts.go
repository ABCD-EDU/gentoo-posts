package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetPostByID(postId string) (*PostSchema, error) {
	post := new(PostSchema)

	sqlQuery := `
		SELECT * FROM posts
		WHERE post_id=$1
		LIMIT 1;
		`

	rows, err := db.Query(sqlQuery, postId)
	if err != nil {
		log.Fatal(err)
		return post, err
	}

	for rows.Next() {
		var postId, userId string
		var content string
		var createdOn time.Time

		err := rows.Scan(&postId, &userId, &content, &createdOn)
		if err != nil {
			log.Fatal(err)
			return post, err
		}

		postInfo := &Post{userId, content, createdOn}
		post = &PostSchema{postId, *postInfo}
	}

	return post, nil
}

func GetPostsFromUser(userId string) ([]PostSchema, error) {
	posts := []PostSchema{}

	sqlQuery := `
	SELECT * FROM posts
	WHERE user_id=$1
	ORDER BY created_on DESC;
	`

	rows, err := db.Query(
		sqlQuery,
		userId,
	)
	if err != nil {
		fmt.Println(err)
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		var postId, userId string
		var content string
		var createdOn time.Time

		err := rows.Scan(&postId, &userId, &content, &createdOn)
		if err != nil {
			fmt.Println(err)
		}

		post := &Post{userId, content, createdOn}
		postSchema := &PostSchema{PostId: postId, PostInfo: *post}
		posts = append(posts, *postSchema)
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
