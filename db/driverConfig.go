package dbConfig

import "fmt"

type Article struct {
	ID    int
	Title string
	Body  []byte
}

const PostgresDriver = "postgres"
const User = "teste"
const Host = "172.17.0.2" // Use o endereço IP do contêiner PostgreSQL
const Port = "5432"
const Password = "teste"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s sslmode=disable", Host, Port, User, Password)
