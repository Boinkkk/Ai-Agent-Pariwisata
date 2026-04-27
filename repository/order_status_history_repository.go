package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderStatusHistoryRepositoryInterface interface {
	Insert(ctx context.Context, history *models.OrderStatusHistory) error
	FindByID(ctx context.Context, id string) (*models.OrderStatusHistory, error)
	FindAll(ctx context.Context) ([]models.OrderStatusHistory, error)
	Update(ctx context.Context, id string, history *models.OrderStatusHistory) error
	Delete(ctx context.Context, id string) error
}

type OrderStatusHistoryRepository struct {
	db *pgxpool.Pool
}

func NewOrderStatusHistoryRepository(db *pgxpool.Pool) *OrderStatusHistoryRepository {
	return &OrderStatusHistoryRepository{db: db}
}

func (r *OrderStatusHistoryRepository) Insert(ctx context.Context, history *models.OrderStatusHistory) error {
	query := `INSERT INTO orders.order_status_history (id, order_id, status, notes, created_at) VALUES ($1, $2, $3, $4, NOW())`
	history.ID = uuid.NewString()
	_, err := r.db.Exec(ctx, query, history.ID, history.OrderID, history.Status, nullIfEmpty(history.Notes))
	return err
}

func (r *OrderStatusHistoryRepository) FindByID(ctx context.Context, id string) (*models.OrderStatusHistory, error) {
	query := `SELECT id::text, order_id::text, status, COALESCE(notes, ''), created_at::text FROM orders.order_status_history WHERE id = $1`
	var history models.OrderStatusHistory
	err := r.db.QueryRow(ctx, query, id).Scan(&history.ID, &history.OrderID, &history.Status, &history.Notes, &history.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &history, nil
}

func (r *OrderStatusHistoryRepository) FindAll(ctx context.Context) ([]models.OrderStatusHistory, error) {
	query := `SELECT id::text, order_id::text, status, COALESCE(notes, ''), created_at::text FROM orders.order_status_history ORDER BY created_at DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	histories := []models.OrderStatusHistory{}
	for rows.Next() {
		var history models.OrderStatusHistory
		if err := rows.Scan(&history.ID, &history.OrderID, &history.Status, &history.Notes, &history.CreatedAt); err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, rows.Err()
}

func (r *OrderStatusHistoryRepository) Update(ctx context.Context, id string, history *models.OrderStatusHistory) error {
	query := `UPDATE orders.order_status_history SET order_id = $2, status = $3, notes = $4 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, history.OrderID, history.Status, nullIfEmpty(history.Notes))
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *OrderStatusHistoryRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders.order_status_history WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
