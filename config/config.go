package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB // the supported drivers are postgresql,mysql and which use sql relational databases

func ConnectDB() {

	// to connect to the databse sql requires username,password and host right either we can create seperate vairables
	// or we use datasource with  a single string

	dsn := "user=postgres password=Helper@123 dbname=postgres sslmode=disable"
	// dsn := "user=postgres password=Helper@123 database=postgres sslmode=disable port=xxxx host=localhost"
	// if we don't mention host and port it will default like 5432 port and host as localhost

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("Error opening databse", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error connecting to database", err)
		panic(err)
	}
	DB = db
	fmt.Println("Successfully connected to database...")

}
