package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	database, err := sql.Open("postgres", "user=postgres password=Pass1234 host=127.0.0.1 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer database.Close()

	stagement, err := database.Prepare("SELECT id, name, age FROM users where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}
	rowId := 1
	row := stagement.QueryRow(rowId)
	var id, age int
	var name string
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}
	fmt.Println("one row", id, name, age)
}
