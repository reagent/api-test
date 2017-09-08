package main

import (
	"os"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	databaseUrl, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatal("$DATABASE_URL must be set")
	}

	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalf("Database connection failed: %s\n", err.Error())
	}

	_, err = db.Exec(`INSERT INTO counter (created_at) VALUES (NOW())`)

	if err != nil {
		log.Fatalf("Insert failed: %s\n", err.Error())
	}
}