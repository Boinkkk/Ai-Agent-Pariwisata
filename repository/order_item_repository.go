package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderItemRepositoryInterface interface {
	Insert(ctx context.Context, item *models.OrderItem) error
	FindByID(ctx context.Context, id string) (*models.OrderItem, error)
	FindAll(ctx context.Context) ([]models.OrderItem, error)
	Update(ctx context.Context, id string, item *models.OrderItem) error
	Delete(ctx context.Context, id string) error
}

type OrderItemRepository struct {
	db *pgxpool.Pool
}

func NewOrderItemRepository(db *pgxpool.Pool) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (r *OrderItemRepository) Insert(ctx context.Context, item *models.OrderItem) error {
	query := `INSERT INTO orders.order_items (id, order_id, product_id, quantity, price_at_purchase) VALUES ($1, $2, $3, $4, $5)`
	item.ID = uuid.New()
	_, err := r.db.Exec(ctx, query, item.ID, item.OrderID, item.ProductID, item.Quantity, item.PriceAtPurchase)
	return err
}

func (r *OrderItemRepository) FindByID(ctx context.Context, id string) (*models.OrderItem, error) {
	query := `SELECT id, order_id, product_id, quantity, price_at_purchase FROM orders.order_items WHERE id = $1`
	var item models.OrderItem
	err := r.db.QueryRow(ctx, query, id).Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.PriceAtPurchase)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *OrderItemRepository) FindAll(ctx context.Context) ([]models.OrderItem, error) {
	query := `SELECT id, order_id, product_id, quantity, price_at_purchase FROM orders.order_items ORDER BY id DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.OrderItem{}
	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.PriceAtPurchase); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *OrderItemRepository) Update(ctx context.Context, id string, item *models.OrderItem) error {
	query := `UPDATE orders.order_items SET order_id = $2, product_id = $3, quantity = $4, price_at_purchase = $5 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, item.OrderID, item.ProductID, item.Quantity, item.PriceAtPurchase)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *OrderItemRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders.order_items WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
