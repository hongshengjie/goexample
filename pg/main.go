package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM example.all_type ")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int64
		var str string
		var str1 string
		var inet string
		rows.Scan(&id, &str, &str1, &inet)
		fmt.Println(id, str, str1, inet)
	}

}
