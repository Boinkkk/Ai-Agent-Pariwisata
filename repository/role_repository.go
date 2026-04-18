package repository

import (
	"context"
	"tutorial/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RoleRepository struct {
	db *pgxpool.Pool
}

func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) GetRoles(ctx context.Context) ([]models.Role, error) {
	query := `SELECT id, name from auth.roles`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var roles []models.Role

	for rows.Next() {
		var role models.Role

		err := rows.Scan(
			&role.ID,
			&role.Name,
		)

		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil

}
