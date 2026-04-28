package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepositoryInterface interface {
	Insert(ctx context.Context, product *models.Product) error
	FindByID(ctx context.Context, id string) (*models.Product, error)
	FindAll(ctx context.Context) ([]models.Product, error)
	Update(ctx context.Context, id string, product *models.Product) error
	Delete(ctx context.Context, id string) error
	GetAllProduct(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	GetProductBySlug(ctx context.Context, slug string) (*models.Product, error)
	AddProduct(ctx context.Context, product *models.Product) error
	UpdateProductByID(ctx context.Context, id string, product models.Product) error
	DeleteProductByID(ctx context.Context, id string) error
}

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) scanProduct(row pgx.Row) (*models.Product, error) {
	var product models.Product
	err := row.Scan(
		&product.ID,
		&product.CategoryID,
		&product.Name,
		&product.Slug,
		&product.Description,
		&product.Price,
		&product.StockQuantity,
		&product.WeightGrams,
		&product.ImageURL,
		&product.IsActive,
		&product.AverageRating,
		&product.Benefit,
		&product.Composition,
		&product.Directions,
		&product.StorageInstruction,
		&product.Manufacturer,
		&product.MarketingLocation,
		&product.ProductionLocation,
		&product.Regency,
		&product.Licensing,
		&product.LicensingNumber,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func productSelectQuery() string {
	return `SELECT id, category_id, name, slug, description, price, stock_quantity, weight_grams, image_url, is_active, average_rating, benefit, composition, directions, storage_instructions, manufacturer, marketing_location, production_location, regency, licensing, licensing_number, created_at, updated_at FROM catalog.products`
}

func (r *ProductRepository) Insert(ctx context.Context, product *models.Product) error {
	query := `INSERT INTO catalog.products (id, category_id, name, slug, description, price, stock_quantity, weight_grams, image_url, is_active, average_rating, benefit, composition, directions, storage_instructions, manufacturer, marketing_location, production_location, regency, licensing, licensing_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, NOW(), NOW())`

	product.ID = uuid.New()

	_, err := r.db.Exec(ctx, query,
		product.ID,
		product.CategoryID,
		product.Name,
		product.Slug,
		product.Description,
		product.Price,
		product.StockQuantity,
		product.WeightGrams,
		product.ImageURL,
		product.IsActive,
		product.AverageRating,
		product.Benefit,
		product.Composition,
		product.Directions,
		product.StorageInstruction,
		product.Manufacturer,
		product.MarketingLocation,
		product.ProductionLocation,
		product.Regency,
		product.Licensing,
		product.LicensingNumber,
	)

	return err
}

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*models.Product, error) {
	query := productSelectQuery() + ` WHERE id = $1`
	return r.scanProduct(r.db.QueryRow(ctx, query, id))
}

func (r *ProductRepository) FindBySlug(ctx context.Context, slug string) (*models.Product, error) {
	query := productSelectQuery() + ` WHERE slug = $1`
	return r.scanProduct(r.db.QueryRow(ctx, query, slug))
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]models.Product, error) {
	query := productSelectQuery() + ` ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.CategoryID,
			&product.Name,
			&product.Slug,
			&product.Description,
			&product.Price,
			&product.StockQuantity,
			&product.WeightGrams,
			&product.ImageURL,
			&product.IsActive,
			&product.AverageRating,
			&product.Benefit,
			&product.Composition,
			&product.Directions,
			&product.StorageInstruction,
			&product.Manufacturer,
			&product.MarketingLocation,
			&product.ProductionLocation,
			&product.Regency,
			&product.Licensing,
			&product.LicensingNumber,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, rows.Err()
}

func (r *ProductRepository) Update(ctx context.Context, id string, product *models.Product) error {
	query := `UPDATE catalog.products SET category_id = $2, name = $3, slug = $4, description = $5, price = $6, stock_quantity = $7, weight_grams = $8, image_url = $9, is_active = $10, average_rating = $11, benefit = $12, composition = $13, directions = $14, storage_instructions = $15, manufacturer = $16, marketing_location = $17, production_location = $18, regency = $19, licensing = $20, licensing_number = $21, updated_at = NOW() WHERE id = $1`

	tag, err := r.db.Exec(ctx, query,
		id,
		product.CategoryID,
		product.Name,
		product.Slug,
		product.Description,
		product.Price,
		product.StockQuantity,
		product.WeightGrams,
		product.ImageURL,
		product.IsActive,
		product.AverageRating,
		product.Benefit,
		product.Composition,
		product.Directions,
		product.StorageInstruction,
		product.Manufacturer,
		product.MarketingLocation,
		product.ProductionLocation,
		product.Regency,
		product.Licensing,
		product.LicensingNumber,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM catalog.products WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *ProductRepository) GetAllProduct(ctx context.Context) ([]models.Product, error) {
	return r.FindAll(ctx)
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	return r.FindByID(ctx, id)
}

func (r *ProductRepository) GetProductBySlug(ctx context.Context, slug string) (*models.Product, error) {
	return r.FindBySlug(ctx, slug)
}

func (r *ProductRepository) AddProduct(ctx context.Context, product *models.Product) error {
	return r.Insert(ctx, product)
}

func (r *ProductRepository) UpdateProductByID(ctx context.Context, id string, product models.Product) error {
	return r.Update(ctx, id, &product)
}

func (r *ProductRepository) DeleteProductByID(ctx context.Context, id string) error {
	return r.Delete(ctx, id)
}
