package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@db:5432?sslmode=disable")
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}
	defer conn.Close(ctx)

	if err = conn.Ping(ctx); err != nil {
		log.Fatalf("can't ping: %v", err)
	}
	fmt.Print("Database is connected\n")

	http.HandleFunc("/", rootHandler)

	fmt.Print("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
