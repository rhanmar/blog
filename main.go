package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

var conn *pgx.Conn

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := getArticles(r.Context())
	if err != nil {
		log.Fatalf("%v", err)
	}

	renderJSON(w, articles)
}

func getArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/article/"):])
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("Received ID: %d \n", id)
	article, err := getArticleByID(r.Context(), int64(id))
	if err != nil {
		log.Fatalf("%v", err)
	}
	if article == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	renderJSON(w, article)
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type Article struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func getArticles(ctx context.Context) ([]Article, error) {
	qb := squirrel.Select("*").From("article")
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query(ctx, sql, args...)
	defer rows.Close()
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	articles := make([]Article, 0)

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.Text, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func getArticleByID(ctx context.Context, id int64) (*Article, error) {
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("*").From("article").Where(squirrel.Eq{"id": id})
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	row := conn.QueryRow(ctx, sql, args...)
	var article Article
	err = row.Scan(&article.ID, &article.Title, &article.Text, &article.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		fmt.Println("No article found")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func main() {
	ctx := context.Background()
	var err error
	conn, err = pgx.Connect(ctx, "postgres://postgres:postgres@db:5432?sslmode=disable")
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}
	defer conn.Close(ctx)

	if err = conn.Ping(ctx); err != nil {
		log.Fatalf("can't ping: %v", err)
	}
	fmt.Print("Database is connected\n")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/articles/", getArticlesHandler)
	http.HandleFunc("/article/", getArticleByIDHandler)

	fmt.Print("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
