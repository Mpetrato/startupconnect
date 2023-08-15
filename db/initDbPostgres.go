package dbConfig

import (
	"database/sql"
	"fmt"
	"time"
)

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

var Db *sql.DB
var err error

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func InitDbPostgres() {
	Db, err = sql.Open(PostgresDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
}
