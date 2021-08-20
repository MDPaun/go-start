package main

import (
	"database/sql"
	"fmt"
	"log"

	//for connecting to db
	client "github.com/MDPaun/go-start/tree/main/website/internal/database/postgres"
	models "github.com/MDPaun/go-start/website/pkg/models/user"
	// _ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432 //5050
	user     = "root"
	password = "root"
	dbname   = "db_main"
)

// var (
// 	DbClient *sql.DB
// )

func main() {
	// connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// var err error
	// DbClient, err := sql.Open("postgres", connectionString)
	// if err != nil {
	// 	panic(err)
	// }
	// defer DbClient.Close()

	// err = DbClient.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Successfully connected!")

	// fmt.Println(DbClient.QueryRow("SELECT * FROM staff WHERE firstname = $1", "Marius"))
	db := client.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM staff;")
	handlerows(rows, err)

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }

	// http.HandleFunc("/", helloHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := []models.User{}
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.GroupID, &u.Email, &u.Password, &u.Salt, &u.FirstName, &u.LastName, &u.Image, &u.IP, &u.Status, &u.DateAdded)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
	fmt.Printf("%T", users)
	// return users
}
