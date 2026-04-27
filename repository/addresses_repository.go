package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AddressesRepositoryInterface interface {
	Insert(ctx context.Context, address *models.Addresses) error
	FindByID(ctx context.Context, id string) (*models.Addresses, error)
	FindAll(ctx context.Context) ([]models.Addresses, error)
	Update(ctx context.Context, id string, address *models.Addresses) error
	Delete(ctx context.Context, id string) error
}

type AddressesRepository struct {
	db *pgxpool.Pool
}

func NewAddressesRepository(db *pgxpool.Pool) *AddressesRepository {
	return &AddressesRepository{db: db}
}

func (r *AddressesRepository) Insert(ctx context.Context, address *models.Addresses) error {
	query := `INSERT INTO orders.addresses (id, user_id, label, recipient_name, phone_number, address_line, city, province, postal_code, is_main) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	address.ID = uuid.New()
	_, err := r.db.Exec(ctx, query, address.ID, address.UserID, address.Label, address.RecipientName, address.PhoneNumber, address.AddressLine, address.City, address.Province, address.PostalCode, address.IsMain)
	return err
}

func (r *AddressesRepository) FindByID(ctx context.Context, id string) (*models.Addresses, error) {
	query := `SELECT id, user_id, label, recipient_name, phone_number, address_line, city, province, postal_code, is_main FROM orders.addresses WHERE id = $1`
	var address models.Addresses
	err := r.db.QueryRow(ctx, query, id).Scan(&address.ID, &address.UserID, &address.Label, &address.RecipientName, &address.PhoneNumber, &address.AddressLine, &address.City, &address.Province, &address.PostalCode, &address.IsMain)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *AddressesRepository) FindAll(ctx context.Context) ([]models.Addresses, error) {
	query := `SELECT id, user_id, label, recipient_name, phone_number, address_line, city, province, postal_code, is_main FROM orders.addresses ORDER BY label ASC`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	addresses := []models.Addresses{}
	for rows.Next() {
		var address models.Addresses
		if err := rows.Scan(&address.ID, &address.UserID, &address.Label, &address.RecipientName, &address.PhoneNumber, &address.AddressLine, &address.City, &address.Province, &address.PostalCode, &address.IsMain); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, rows.Err()
}

func (r *AddressesRepository) Update(ctx context.Context, id string, address *models.Addresses) error {
	query := `UPDATE orders.addresses SET user_id = $2, label = $3, recipient_name = $4, phone_number = $5, address_line = $6, city = $7, province = $8, postal_code = $9, is_main = $10 WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id, address.UserID, address.Label, address.RecipientName, address.PhoneNumber, address.AddressLine, address.City, address.Province, address.PostalCode, address.IsMain)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *AddressesRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders.addresses WHERE id = $1`
	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
