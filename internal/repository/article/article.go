package article

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	articleSvc "github.com/rhanmar/blog/internal/services/article"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{conn: conn}
}

func (r *Repository) GetArticles(ctx context.Context) ([]*articleSvc.Article, error) {
	qb := squirrel.Select("*").From("article")
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.conn.Query(ctx, sql, args...)
	defer rows.Close()
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	articles := make([]*Article, 0)

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.Text, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}
	return toServiceArticles(articles), nil
}

func (r *Repository) GetArticleByID(ctx context.Context, id int64) (*articleSvc.Article, error) {
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).Select("*").From("article").Where(squirrel.Eq{"id": id})
	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.conn.QueryRow(ctx, sql, args...)
	var article Article
	err = row.Scan(&article.ID, &article.Title, &article.Text, &article.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		fmt.Println("No article found")
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return toServiceArticle(article), nil
}
