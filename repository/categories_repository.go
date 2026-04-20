package repository

import (
	"context"
	"tutorial/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesRepository struct {
	db *pgxpool.Pool
}

func NewCategoriesRepository(db *pgxpool.Pool) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

func (r *CategoriesRepository) GetCategories(ctx context.Context) ([]models.Categories, error) {
	query := `SELECT id, name, slug, description FROM catalog.categories`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []models.Categories

	for rows.Next() {
		var categorie models.Categories

		err := rows.Scan(
			&categorie.ID,
			&categorie.Name,
			&categorie.Slug,
			&categorie.Description,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, categorie)
	}

	return categories, nil
}
