package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
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
