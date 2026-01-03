package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nill {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nill {
		log.Fatal("Database Is Not Reachable")
	}

	log.Println("Database Connected")

}
