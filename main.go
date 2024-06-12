package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rhanmar/blog/internal/handlers"
	articleRepo "github.com/rhanmar/blog/internal/repository/article"
	articleSvc "github.com/rhanmar/blog/internal/services/article"
	"log"
	"net/http"
)

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

	articleRepository := articleRepo.NewRepository(conn)
	articleService := articleSvc.NewService(articleRepository)

	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/articles/", handlers.GetArticles(articleService))
	http.HandleFunc("/article/", handlers.GetArticleByID(articleService))

	fmt.Print("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
