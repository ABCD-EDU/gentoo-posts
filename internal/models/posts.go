package models

import (
	"fmt"
	"log"
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
	WHERE user_id=$1;
	`

	rows, err := db.Query(
		sqlQuery,
		userId,
	)
	if err != nil {
		log.Fatal(err)
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		var postId, userId string
		var content string
		var createdOn time.Time

		err := rows.Scan(&postId, &userId, &content, &createdOn)
		if err != nil {
			log.Fatal(err)
			return posts, err
		}

		post := &Post{userId, content, createdOn}
		postSchema := &PostSchema{PostId: postId, PostInfo: *post}
		posts = append(posts, *postSchema)
	}

	return posts, nil
}

func WritePost(post Post) (string, error) {
	sqlQuery := `
	INSERT INTO posts(user_id, content, createdOn)
	VALUES ($1, $2, $3)
	RETURNING post_id;
	`

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

	return postId, nil
}

func WriteReplyPost() {

}
