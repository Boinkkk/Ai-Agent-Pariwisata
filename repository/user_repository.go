package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	Insert(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, id string, user *models.User) error
	Delete(ctx context.Context, id string) error
	CreateUser(ctx context.Context, user *models.User) error
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(ctx context.Context, user *models.User) error {
	query := `INSERT INTO auth.users (id, role_id, username, email, password_hash, full_name, phone_number) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	user.ID = uuid.New()

	_, err := r.db.Exec(ctx, query,
		user.ID,
		user.RoleID,
		user.Username,
		user.Email,
		user.Password,
		user.FullName,
		user.PhoneNumber,
	)

	return err
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, role_id, username, email, password_hash, full_name, phone_number, created_at, updated_at FROM auth.users WHERE id = $1`

	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.RoleID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, role_id, username, email, password_hash, full_name, phone_number, created_at, updated_at FROM auth.users ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.RoleID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.FullName,
			&user.PhoneNumber,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

func (r *UserRepository) Update(ctx context.Context, id string, user *models.User) error {
	query := `UPDATE auth.users SET role_id = $2, username = $3, email = $4, password_hash = COALESCE(NULLIF($5, ''), password_hash), full_name = $6, phone_number = $7, updated_at = NOW() WHERE id = $1`

	tag, err := r.db.Exec(ctx, query,
		id,
		user.RoleID,
		user.Username,
		user.Email,
		user.Password,
		user.FullName,
		user.PhoneNumber,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM auth.users WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.Insert(ctx, user)
}
