package article

import (
	"context"
)

type articleRepository interface {
	GetArticles(context.Context) ([]*Article, error)
	GetArticleByID(context.Context, int64) (*Article, error)
}

type Service struct {
	articleRepo articleRepository
}

func NewService(articleRepo articleRepository) *Service {
	return &Service{
		articleRepo: articleRepo,
	}
}

func (s *Service) GetArticles(ctx context.Context) ([]*Article, error) {
	return s.articleRepo.GetArticles(ctx)
}

func (s *Service) GetArticleByID(ctx context.Context, id int64) (*Article, error) {
	return s.articleRepo.GetArticleByID(ctx, id)
}
