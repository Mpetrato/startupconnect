package dbConfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Article struct {
	ID    int
	Title string
	Body  []byte
}

var (
	PostgresDriver string
	Usert          string
	Host           string
	Port           string
	Password       string
	Database       string
	SslMode        = "disable"
	DataSourceName string
)

func InitDriver() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PostgresDriver = os.Getenv("POSTGRES_DRIVER")
	Usert = os.Getenv("USER")
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	Password = os.Getenv("PASSWORD")
	Database = os.Getenv("DATABASE")

	DataSourceName = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		Usert,
		Password,
		Host,
		Port,
		Database,
		SslMode,
	)
}
