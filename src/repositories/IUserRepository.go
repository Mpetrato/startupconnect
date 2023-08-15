package repository

import (
	"database/sql"

	dbConfig "github.com/Mpetrato/startupconnect/db"
	"github.com/Mpetrato/startupconnect/src/types"
)

// var db *sql.DB
// var err error

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

func Register(user *types.User) error {
	_, err := dbConfig.Db.Exec("INSERT INTO users (id, username, email, password) VALUES ($1 ,$2, $3, $4)",
		user.Id, user.Username, user.Email, user.Password)
	return err
}

func GetUserByEmail(email string) (*types.User, error) {
	query := "SELECT * FROM users WHERE email = $1"

	row := dbConfig.Db.QueryRow(query, email)

	var user types.User
	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
