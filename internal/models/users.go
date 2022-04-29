package models

import "fmt"

func WriteUserRegistration(user UserSchema) (string, error) {
	sqlQuery := `
			INSERT INTO users(user_id, username, email, google_photo)
			VALUES($1,$2,$3,$4)
		`

	fmt.Printf("id: %s\nusername: %s\nemail: %s\nphoto: %s\n",
		user.UserId,
		user.UserInfo.Username,
		user.UserInfo.Email,
		user.UserInfo.Photo,
	)

	_, err := db.Exec(
		sqlQuery,
		user.UserId,
		user.UserInfo.Username,
		user.UserInfo.Email,
		user.UserInfo.Photo,
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR WITH WRITING DATA TO DATABASE")
		return "", err
	}

	return user.UserId, nil
}
