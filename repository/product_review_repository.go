package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductReviewRepositoryInterface interface {
	Insert(ctx context.Context, review *models.ProductReview) error
	FindByID(ctx context.Context, id string) (*models.ProductReview, error)
	FindAll(ctx context.Context) ([]models.ProductReview, error)
	Update(ctx context.Context, id string, review *models.ProductReview) error
	Delete(ctx context.Context, id string) error
}

type ProductReviewRepository struct {
	db *pgxpool.Pool
}

func NewProductReviewRepository(db *pgxpool.Pool) *ProductReviewRepository {
	return &ProductReviewRepository{db: db}
}

func (r *ProductReviewRepository) Insert(ctx context.Context, review *models.ProductReview) error {
	query := `INSERT INTO catalog.product_reviews (id, product_id, user_id, rating, comment, created_at) VALUES ($1, $2, $3, $4, $5, NOW())`
	review.ID = uuid.New()
	_, err := r.db.Exec(ctx, query, review.ID, review.ProductID, review.UserID, review.Rating, review.Comment)
	return err
}

func (r *ProductReviewRepository) FindByID(ctx context.Context, id string) (*models.ProductReview, error) {
	query := `SELECT id, product_id, user_id, rating, comment, created_at FROM catalog.product_reviews WHERE id = $1`
	var review models.ProductReview
	err := r.db.QueryRow(ctx, query, id).Scan(&review.ID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &review.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ProductReviewRepository) FindAll(ctx context.Context) ([]models.ProductReview, error) {
	query := `SELECT id, product_id, user_id, rating, comment, created_at FROM catalog.product_reviews ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	reviews := []models.ProductReview{}
	for rows.Next() {
		var review models.ProductReview
		if err := rows.Scan(&review.ID, &review.ProductID, &review.UserID, &review.Rating, &review.Comment, &review.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, rows.Err()
}

func (r *ProductReviewRepository) Update(ctx context.Context, id string, review *models.ProductReview) error {
	query := `UPDATE catalog.product_reviews SET product_id = $2, user_id = $3, rating = $4, comment = $5 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, review.ProductID, review.UserID, review.Rating, review.Comment)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ProductReviewRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM catalog.product_reviews WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
