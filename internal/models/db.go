package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	db *sql.DB
)

func InitializeDB() {
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	dbname := viper.GetString("db.dbname")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)
	var err error
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
	}
}
