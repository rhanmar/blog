package handlers

import (
	"fmt"
	articleSvc "github.com/rhanmar/blog/internal/services/article"
	"log"
	"net/http"
	"strconv"
)

func GetArticles(articleService *articleSvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := articleService.GetArticles(r.Context())
		if err != nil {
			log.Fatalf("[GetArticles] %v", err)
		}
		renderJSON(w, articles)
	}
}

func GetArticleByID(articleService *articleSvc.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Path[len("/article/"):])
		if err != nil {
			log.Fatalf("[GetArticleByID] %v", err)
		}
		fmt.Printf("Received ID: %d \n", id)
		article, err := articleService.GetArticleByID(r.Context(), int64(id))
		if err != nil {
			log.Fatalf("%v", err)
		}
		if article == nil {
			w.WriteHeader(http.StatusNotFound)
		}
		renderJSON(w, article)
	}
}
