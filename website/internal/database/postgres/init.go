package postgres

import (
	"database/sql"
	"fmt"

	//for connecting to db
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432 //5050
	user     = "root"
	password = "root"
	dbname   = "db_main"
)

var (
	DbClient *sql.DB
)

func init() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DbClient, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer DbClient.Close()

	err = DbClient.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}

// func main() {

// }
