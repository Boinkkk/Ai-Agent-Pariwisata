package repository

import (
	"context"
	"tutorial/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetPasswordByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password_hash AS password FROM auth.users WHERE email = $1`

	var user models.User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	// fmt.Println(user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
