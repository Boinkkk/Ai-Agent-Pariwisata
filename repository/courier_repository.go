package repository

import (
	"context"
	"tutorial/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CourierRepositoryInterface interface {
	Insert(ctx context.Context, courier *models.Courier) error
	FindByID(ctx context.Context, id int) (*models.Courier, error)
	FindAll(ctx context.Context) ([]models.Courier, error)
	Update(ctx context.Context, id int, courier *models.Courier) error
	Delete(ctx context.Context, id int) error
}

type CourierRepository struct {
	db *pgxpool.Pool
}

func NewCourierRepository(db *pgxpool.Pool) *CourierRepository {
	return &CourierRepository{db: db}
}

func (r *CourierRepository) Insert(ctx context.Context, courier *models.Courier) error {
	query := `INSERT INTO orders.couriers (code, name, service_type) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(ctx, query, courier.Code, courier.Name, courier.ServiceType).Scan(&courier.ID)
}

func (r *CourierRepository) FindByID(ctx context.Context, id int) (*models.Courier, error) {
	query := `SELECT id, code, name, service_type FROM orders.couriers WHERE id = $1`
	var courier models.Courier
	err := r.db.QueryRow(ctx, query, id).Scan(&courier.ID, &courier.Code, &courier.Name, &courier.ServiceType)
	if err != nil {
		return nil, err
	}
	return &courier, nil
}

func (r *CourierRepository) FindAll(ctx context.Context) ([]models.Courier, error) {
	query := `SELECT id, code, name, service_type FROM orders.couriers ORDER BY id ASC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	couriers := []models.Courier{}
	for rows.Next() {
		var courier models.Courier
		if err := rows.Scan(&courier.ID, &courier.Code, &courier.Name, &courier.ServiceType); err != nil {
			return nil, err
		}
		couriers = append(couriers, courier)
	}
	return couriers, rows.Err()
}

func (r *CourierRepository) Update(ctx context.Context, id int, courier *models.Courier) error {
	query := `UPDATE orders.couriers SET code = $2, name = $3, service_type = $4 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, courier.Code, courier.Name, courier.ServiceType)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *CourierRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM orders.couriers WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
