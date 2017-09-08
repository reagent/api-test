package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		log.Fatal("$PORT must be set")
	} 

	databaseUrl, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatal("$DATABASE_URL must be set")
	}

	db, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalf("Database connection failed: %s\n", err.Error())
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS counter (created_at TIMESTAMP NOT NULL)`)

	if err != nil {
		log.Fatalf("Schema creation failed: %s\n", err.Error())
	}

	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var count int64

		w.Header().Set("Content-Type", "application/json")

		row := db.QueryRow(`SELECT COUNT(*) FROM counter`)
		if err := row.Scan(&count); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, fmt.Sprintf(`{"msg":"%s"}` + "\n", err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, fmt.Sprintf(`{"count":"%d"}` + "\n", count))
		}
	})

	err = http.ListenAndServe(":"+port, handler)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
