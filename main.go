package main

import (
	"database/sql"
	"fmt"
	"time"

	dbConfig "github.com/Mpetrato/startupconnect/db"
	"github.com/Mpetrato/startupconnect/router"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {
	router.InitRouter()

	fmt.Printf("Accessing... ")

	db, err = sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	sqlSelect()
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func sqlSelect() {
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
		checkErr(err)
		users = append(users, user)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Username: %s, Email: %s, CreatedAt: %s\n", user.ID, user.Username, user.Email, user.CreatedAt)
	}
}
