-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS article (
    id BIGINT  PRIMARY KEY,
    title TEXT NOT NULL,
    text TEXT NULL,
--     created_at TIMESTAMP WITH TIME ZONE NOT NULL
    created_at TIMESTAMP DEFAULT NOW() NOT NULL
);

COMMENT ON COLUMN article.id IS 'ID Статьи';
COMMENT ON COLUMN article.title IS 'Заголовок Статьи';
COMMENT ON COLUMN article.text IS 'Текст Статьи';
COMMENT ON COLUMN article.created_at IS 'Дата создания Статьи';
COMMENT ON TABLE article IS 'Статья';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS article;
-- +goose StatementEnd
