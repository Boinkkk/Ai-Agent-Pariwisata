package repository

import (
	"context"
	"tutorial/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesRepositoryInterface interface {
	Insert(ctx context.Context, category *models.Categories) error
	FindByID(ctx context.Context, id int) (*models.Categories, error)
	FindAll(ctx context.Context) ([]models.Categories, error)
	Update(ctx context.Context, id int, category *models.Categories) error
	Delete(ctx context.Context, id int) error
	GetCategories(ctx context.Context) ([]models.Categories, error)
	AddCategorie(ctx context.Context, category *models.Categories) error
	DeleteCategoriesByID(ctx context.Context, id int) error
}

type CategoriesRepository struct {
	db *pgxpool.Pool
}

func NewCategoriesRepository(db *pgxpool.Pool) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

func (r *CategoriesRepository) Insert(ctx context.Context, category *models.Categories) error {
	query := `INSERT INTO catalog.categories (name, slug, description) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(ctx, query, category.Name, category.Slug, category.Description).Scan(&category.ID)
}

func (r *CategoriesRepository) FindByID(ctx context.Context, id int) (*models.Categories, error) {
	query := `SELECT id, name, slug, description FROM catalog.categories WHERE id = $1`

	var category models.Categories
	err := r.db.QueryRow(ctx, query, id).Scan(&category.ID, &category.Name, &category.Slug, &category.Description)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoriesRepository) FindAll(ctx context.Context) ([]models.Categories, error) {
	query := `SELECT id, name, slug, description FROM catalog.categories ORDER BY id ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []models.Categories{}
	for rows.Next() {
		var category models.Categories
		if err := rows.Scan(&category.ID, &category.Name, &category.Slug, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, rows.Err()
}

func (r *CategoriesRepository) Update(ctx context.Context, id int, category *models.Categories) error {
	query := `UPDATE catalog.categories SET name = $2, slug = $3, description = $4 WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id, category.Name, category.Slug, category.Description)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *CategoriesRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM catalog.categories WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *CategoriesRepository) GetCategories(ctx context.Context) ([]models.Categories, error) {
	return r.FindAll(ctx)
}

func (r *CategoriesRepository) AddCategorie(ctx context.Context, category *models.Categories) error {
	return r.Insert(ctx, category)
}

func (r *CategoriesRepository) DeleteCategoriesByID(ctx context.Context, id int) error {
	return r.Delete(ctx, id)
}
