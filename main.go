package main

import (
	dbConfig "github.com/Mpetrato/startupconnect/db"
	"github.com/Mpetrato/startupconnect/src/router"
	_ "github.com/lib/pq"
)

func main() {
	dbConfig.InitDbPostgres()
	router.InitRouter()
}
