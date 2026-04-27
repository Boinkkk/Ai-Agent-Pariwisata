package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartItemRepositoryInterface interface {
	Insert(ctx context.Context, item *models.CartItem) error
	FindByID(ctx context.Context, id string) (*models.CartItem, error)
	FindAll(ctx context.Context) ([]models.CartItem, error)
	Update(ctx context.Context, id string, item *models.CartItem) error
	Delete(ctx context.Context, id string) error
}

type CartItemRepository struct {
	db *pgxpool.Pool
}

func NewCartItemRepository(db *pgxpool.Pool) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (r *CartItemRepository) Insert(ctx context.Context, item *models.CartItem) error {
	query := `INSERT INTO orders.cart_items (id, cart_id, product_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())`
	item.ID = uuid.New()
	_, err := r.db.Exec(ctx, query, item.ID, item.CartID, item.ProductID, item.Quantity)
	return err
}

func (r *CartItemRepository) FindByID(ctx context.Context, id string) (*models.CartItem, error) {
	query := `SELECT id, cart_id, product_id, quantity, created_at, updated_at FROM orders.cart_items WHERE id = $1`
	var item models.CartItem
	err := r.db.QueryRow(ctx, query, id).Scan(&item.ID, &item.CartID, &item.ProductID, &item.Quantity, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CartItemRepository) FindAll(ctx context.Context) ([]models.CartItem, error) {
	query := `SELECT id, cart_id, product_id, quantity, created_at, updated_at FROM orders.cart_items ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.CartItem{}
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ID, &item.CartID, &item.ProductID, &item.Quantity, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *CartItemRepository) Update(ctx context.Context, id string, item *models.CartItem) error {
	query := `UPDATE orders.cart_items SET cart_id = $2, product_id = $3, quantity = $4, updated_at = NOW() WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, item.CartID, item.ProductID, item.Quantity)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *CartItemRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders.cart_items WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
