package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepositoryInterface interface {
	Insert(ctx context.Context, payment *models.Payment) error
	FindByID(ctx context.Context, id string) (*models.Payment, error)
	FindAll(ctx context.Context) ([]models.Payment, error)
	Update(ctx context.Context, id string, payment *models.Payment) error
	Delete(ctx context.Context, id string) error
}

type PaymentRepository struct {
	db *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Insert(ctx context.Context, payment *models.Payment) error {
	query := `INSERT INTO payments.payments (id, external_transaction_id, payment_method, amount, payment_url, paid_at) VALUES ($1, $2, $3, $4, $5, NULLIF($6, '')::timestamptz)`
	payment.ID = uuid.NewString()
	_, err := r.db.Exec(ctx, query, payment.ID, nullIfEmpty(payment.ExternalTransactionID), payment.PaymentMethod, payment.Amount, nullIfEmpty(payment.PaymentUrl), payment.PaidAt)
	return err
}

func (r *PaymentRepository) FindByID(ctx context.Context, id string) (*models.Payment, error) {
	query := `SELECT id::text, COALESCE(external_transaction_id, ''), payment_method, amount, COALESCE(payment_url, ''), COALESCE(paid_at::text, '') FROM payments.payments WHERE id = $1`
	var payment models.Payment
	err := r.db.QueryRow(ctx, query, id).Scan(&payment.ID, &payment.ExternalTransactionID, &payment.PaymentMethod, &payment.Amount, &payment.PaymentUrl, &payment.PaidAt)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) FindAll(ctx context.Context) ([]models.Payment, error) {
	query := `SELECT id::text, COALESCE(external_transaction_id, ''), payment_method, amount, COALESCE(payment_url, ''), COALESCE(paid_at::text, '') FROM payments.payments ORDER BY id DESC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payments := []models.Payment{}
	for rows.Next() {
		var payment models.Payment
		if err := rows.Scan(&payment.ID, &payment.ExternalTransactionID, &payment.PaymentMethod, &payment.Amount, &payment.PaymentUrl, &payment.PaidAt); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, rows.Err()
}

func (r *PaymentRepository) Update(ctx context.Context, id string, payment *models.Payment) error {
	query := `UPDATE payments.payments SET external_transaction_id = $2, payment_method = $3, amount = $4, payment_url = $5, paid_at = NULLIF($6, '')::timestamptz WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, nullIfEmpty(payment.ExternalTransactionID), payment.PaymentMethod, payment.Amount, nullIfEmpty(payment.PaymentUrl), payment.PaidAt)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *PaymentRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM payments.payments WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
