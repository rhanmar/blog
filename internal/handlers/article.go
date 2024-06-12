package handlers

import (
	"context"
	"fmt"
	articleSvc "github.com/rhanmar/blog/internal/services/article"
	"log"
	"net/http"
	"strconv"
)

type articleService interface {
	GetArticles(context.Context) ([]*articleSvc.Article, error)
	GetArticleByID(context.Context, int64) (*articleSvc.Article, error)
}

func GetArticles(articleService articleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := articleService.GetArticles(r.Context())
		if err != nil {
			log.Fatalf("[GetArticles] %v", err)
		}
		renderJSON(w, articles)
	}
}

func GetArticleByID(articleService articleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Path[len("/article/"):])
		if err != nil {
			log.Fatalf("[GetArticleByID] %v", err)
		}
		fmt.Printf("Received ID: %d \n", id)
		article, err := articleService.GetArticleByID(r.Context(), int64(id))
		if err != nil {
			log.Fatalf("[GetArticleByID] %v", err)
		}
		if article == nil {
			w.WriteHeader(http.StatusNotFound)
		}
		renderJSON(w, article)
	}
}
