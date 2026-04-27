package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepositoryInterface interface {
	Insert(ctx context.Context, order *models.Order) error
	FindByID(ctx context.Context, id string) (*models.Order, error)
	FindAll(ctx context.Context) ([]models.Order, error)
	Update(ctx context.Context, id string, order *models.Order) error
	Delete(ctx context.Context, id string) error
}

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Insert(ctx context.Context, order *models.Order) error {
	query := `INSERT INTO orders.orders (id, user_id, payment_id, addresses_id, courier_id, total_items_price, shipping_cost, total_amount, status, tracking_number, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW())`
	order.ID = uuid.NewString()
	_, err := r.db.Exec(ctx, query, order.ID, order.UserID, nullIfEmpty(order.PaymentID), nullIfEmpty(order.AddressesID), nullIfEmpty(order.CourierID), order.TotalItemsPrice, order.ShippingCost, order.TotalAmount, order.Status, nullIfEmpty(order.TrackingNumber))
	return err
}

func (r *OrderRepository) FindByID(ctx context.Context, id string) (*models.Order, error) {
	query := `SELECT id::text, user_id::text, COALESCE(payment_id::text, ''), COALESCE(addresses_id::text, ''), COALESCE(courier_id::text, ''), total_items_price, shipping_cost, total_amount, status, COALESCE(tracking_number, ''), created_at::text FROM orders.orders WHERE id = $1`
	var order models.Order
	err := r.db.QueryRow(ctx, query, id).Scan(&order.ID, &order.UserID, &order.PaymentID, &order.AddressesID, &order.CourierID, &order.TotalItemsPrice, &order.ShippingCost, &order.TotalAmount, &order.Status, &order.TrackingNumber, &order.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindAll(ctx context.Context) ([]models.Order, error) {
	query := `SELECT id::text, user_id::text, COALESCE(payment_id::text, ''), COALESCE(addresses_id::text, ''), COALESCE(courier_id::text, ''), total_items_price, shipping_cost, total_amount, status, COALESCE(tracking_number, ''), created_at::text FROM orders.orders ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []models.Order{}
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.PaymentID, &order.AddressesID, &order.CourierID, &order.TotalItemsPrice, &order.ShippingCost, &order.TotalAmount, &order.Status, &order.TrackingNumber, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, rows.Err()
}

func (r *OrderRepository) Update(ctx context.Context, id string, order *models.Order) error {
	query := `UPDATE orders.orders SET user_id = $2, payment_id = $3, addresses_id = $4, courier_id = $5, total_items_price = $6, shipping_cost = $7, total_amount = $8, status = $9, tracking_number = $10 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, order.UserID, nullIfEmpty(order.PaymentID), nullIfEmpty(order.AddressesID), nullIfEmpty(order.CourierID), order.TotalItemsPrice, order.ShippingCost, order.TotalAmount, order.Status, nullIfEmpty(order.TrackingNumber))
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders.orders WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
