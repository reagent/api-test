package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		log.Fatal("$PORT must be set")
	}

	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, `{"status":"ok"}`)
	})

	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
