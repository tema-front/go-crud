package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(dbURL string) *sql.DB {
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal("Can't ping the database", err)
	}

	log.Println("Database connected successfully")
	return conn
}
