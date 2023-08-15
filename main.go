package main

import (
	"log"

	dbConfig "github.com/Mpetrato/startupconnect/db"
	"github.com/Mpetrato/startupconnect/src/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfig.InitDbPostgres()
	router.InitRouter()
}
