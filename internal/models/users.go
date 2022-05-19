package models

import "fmt"

func WriteUserRegistration(user TempUser) (string, error) {
	sqlQuery := `
			INSERT INTO users(user_id, username, email, google_photo, can_post, is_admin)
			VALUES($1,$2,$3,$4,$5,$6)
		`

	fmt.Printf("id: %s\nusername: %s\nemail: %s\nphoto: %s\nisAdmin: %s\n",
		user.UserId,
		user.UserInfo.Username,
		user.UserInfo.Email,
		user.UserInfo.Photo,
		"false",
	)

	_, err := db.Exec(
		sqlQuery,
		user.UserId,
		user.UserInfo.Username,
		user.UserInfo.Email,
		user.UserInfo.Photo,
		true,
		false,
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR WITH WRITING DATA TO DATABASE")
		return "", err
	}

	return user.UserId, nil
}

func MuteUser(userId string) error {
	sqlQuery := `
		UPDATE users
		SET can_post=false
		WHERE user_id=$1
		`

	_, err := db.Exec(sqlQuery, userId)
	if err != nil {
		return err
	}

	return nil
}

func BanUser(userId string) error {
	sqlQuery := `
		DELETE FROM users
		WHERE user_id=$1
		`

	_, err := db.Exec(sqlQuery, userId)
	if err != nil {
		return err
	}

	return nil
}
