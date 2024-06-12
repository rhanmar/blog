package article

import (
	articleSvc "github.com/rhanmar/blog/internal/services/article"
	"time"
)

type Article struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func toServiceArticle(article Article) *articleSvc.Article {
	return &articleSvc.Article{
		ID:        article.ID,
		Title:     article.Title,
		Text:      article.Text,
		CreatedAt: article.CreatedAt,
	}
}

func toServiceArticles(articles []*Article) []*articleSvc.Article {
	result := make([]*articleSvc.Article, 0, len(articles))
	for _, article := range articles {
		result = append(result, toServiceArticle(*article))
	}
	return result
}
