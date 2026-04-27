package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ArticleRepositoryInterface interface {
	Insert(ctx context.Context, article *models.Article) error
	FindByID(ctx context.Context, id string) (*models.Article, error)
	FindAll(ctx context.Context) ([]models.Article, error)
	Update(ctx context.Context, id string, article *models.Article) error
	Delete(ctx context.Context, id string) error
}

type ArticleRepository struct {
	db *pgxpool.Pool
}

func NewArticleRepository(db *pgxpool.Pool) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Insert(ctx context.Context, article *models.Article) error {
	query := `INSERT INTO catalog.articles (id, author_id, title, slug, content, image_url, is_published, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`
	article.ID = uuid.New()
	_, err := r.db.Exec(ctx, query, article.ID, article.AuthorID, article.Title, article.Slug, article.Content, article.ImageURL, article.IsPublished)
	return err
}

func (r *ArticleRepository) FindByID(ctx context.Context, id string) (*models.Article, error) {
	query := `SELECT id, author_id, title, slug, content, image_url, is_published, created_at FROM catalog.articles WHERE id = $1`
	var article models.Article
	err := r.db.QueryRow(ctx, query, id).Scan(&article.ID, &article.AuthorID, &article.Title, &article.Slug, &article.Content, &article.ImageURL, &article.IsPublished, &article.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *ArticleRepository) FindAll(ctx context.Context) ([]models.Article, error) {
	query := `SELECT id, author_id, title, slug, content, image_url, is_published, created_at FROM catalog.articles ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	articles := []models.Article{}
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.AuthorID, &article.Title, &article.Slug, &article.Content, &article.ImageURL, &article.IsPublished, &article.CreatedAt); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, rows.Err()
}

func (r *ArticleRepository) Update(ctx context.Context, id string, article *models.Article) error {
	query := `UPDATE catalog.articles SET author_id = $2, title = $3, slug = $4, content = $5, image_url = $6, is_published = $7 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, article.AuthorID, article.Title, article.Slug, article.Content, article.ImageURL, article.IsPublished)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ArticleRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM catalog.articles WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
