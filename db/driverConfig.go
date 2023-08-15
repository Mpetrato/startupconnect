package dbConfig

import "fmt"

type Article struct {
	ID    int
	Title string
	Body  []byte
}

// const PostgresDriver = "postgres"
// const Usert = "teste"
// const Host = "localhost"
// const Port = "5432"
// const Password = "teste"

// var DataSourceName = "postgres://teste:teste@localhost:5432/postgres?sslmode=disable"

const (
	PostgresDriver = "postgres"
	Usert          = "teste"
	Host           = "localhost"
	Port           = "5432"
	Password       = "teste"
	Database       = "postgres"
	SslMode        = "disable"
)

var DataSourceName = fmt.Sprintf(
	"postgres://%s:%s@%s:%s/%s?sslmode=%s",
	Usert,
	Password,
	Host,
	Port,
	Database,
	SslMode,
)
