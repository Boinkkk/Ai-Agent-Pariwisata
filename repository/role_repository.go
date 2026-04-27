package repository

import (
	"context"
	"tutorial/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoleRepositoryInterface interface {
	Insert(ctx context.Context, role *models.Role) error
	FindByID(ctx context.Context, id int) (*models.Role, error)
	FindAll(ctx context.Context) ([]models.Role, error)
	Update(ctx context.Context, id int, role *models.Role) error
	Delete(ctx context.Context, id int) error
	GetRoles(ctx context.Context) ([]models.Role, error)
}

type RoleRepository struct {
	db *pgxpool.Pool
}

func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Insert(ctx context.Context, role *models.Role) error {
	query := `INSERT INTO auth.roles (name) VALUES ($1) RETURNING id`
	return r.db.QueryRow(ctx, query, role.Name).Scan(&role.ID)
}

func (r *RoleRepository) FindByID(ctx context.Context, id int) (*models.Role, error) {
	query := `SELECT id, name FROM auth.roles WHERE id = $1`

	var role models.Role
	err := r.db.QueryRow(ctx, query, id).Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) FindAll(ctx context.Context) ([]models.Role, error) {
	query := `SELECT id, name FROM auth.roles ORDER BY id ASC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []models.Role{}
	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, rows.Err()
}

func (r *RoleRepository) Update(ctx context.Context, id int, role *models.Role) error {
	query := `UPDATE auth.roles SET name = $2 WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id, role.Name)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *RoleRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM auth.roles WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *RoleRepository) GetRoles(ctx context.Context) ([]models.Role, error) {
	return r.FindAll(ctx)
}
