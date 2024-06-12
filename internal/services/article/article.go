package article

import (
	"context"
	"github.com/rhanmar/blog/internal/repository/article"
)

type Service struct {
	articleRepo *article.Repository
}

func NewService(articleRepo *article.Repository) *Service {
	return &Service{
		articleRepo: articleRepo,
	}
}

func (s *Service) GetArticles(ctx context.Context) ([]article.Article, error) {
	return s.articleRepo.GetArticles(ctx)
}

func (s *Service) GetArticleByID(ctx context.Context, id int64) (*article.Article, error) {
	return s.articleRepo.GetArticleByID(ctx, id)
}
