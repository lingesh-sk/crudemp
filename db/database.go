package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	connStr := "user=postgres dbname=CRUD sslmode=disable password=15522345"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// if err = db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }
}
