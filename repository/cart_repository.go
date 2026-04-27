package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepositoryInterface interface {
	Insert(ctx context.Context, cart *models.Cart) error
	FindByID(ctx context.Context, id string) (*models.Cart, error)
	FindAll(ctx context.Context) ([]models.Cart, error)
	Update(ctx context.Context, id string, cart *models.Cart) error
	Delete(ctx context.Context, id string) error
}

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) Insert(ctx context.Context, cart *models.Cart) error {
	query := `INSERT INTO orders.carts (id, user_id, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())`
	cart.ID = uuid.New()
	_, err := r.db.Exec(ctx, query, cart.ID, cart.UserID)
	return err
}

func (r *CartRepository) FindByID(ctx context.Context, id string) (*models.Cart, error) {
	query := `SELECT id, user_id, created_at, updated_at FROM orders.carts WHERE id = $1`
	var cart models.Cart
	err := r.db.QueryRow(ctx, query, id).Scan(&cart.ID, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) FindAll(ctx context.Context) ([]models.Cart, error) {
	query := `SELECT id, user_id, created_at, updated_at FROM orders.carts ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	carts := []models.Cart{}
	for rows.Next() {
		var cart models.Cart
		if err := rows.Scan(&cart.ID, &cart.UserID, &cart.CreatedAt, &cart.UpdatedAt); err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, rows.Err()
}

func (r *CartRepository) Update(ctx context.Context, id string, cart *models.Cart) error {
	query := `UPDATE orders.carts SET user_id = $2, updated_at = NOW() WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, cart.UserID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *CartRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders.carts WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
