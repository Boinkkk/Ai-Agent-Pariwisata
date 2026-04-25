package repository

import (
	"context"
	"tutorial/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProduct(ctx context.Context) ([]models.Product, error) {
	query := `SELECT 
				id,
				category_id,
				name,
				slug,
				description, 
				price,
				stock_quantity,
				weight_grams,
				image_url,
				is_active,
				average_rating,
				benefit,
				composition,
				directions,
				storage_instructions,
				manufacturer,
				marketing_location,
				production_location,
				regency,
				licensing,
				licensing_number,
				created_at,
				updated_at
			FROM catalog.products;`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product

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

	return products, err

}

func (r *ProductRepository) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	query := `SELECT 
				id,
				category_id,
				name,
				slug,
				description, 
				price,
				stock_quantity,
				weight_grams,
				image_url,
				is_active,
				average_rating,
				benefit,
				composition,
				directions,
				storage_instructions,
				manufacturer,
				marketing_location,
				production_location,
				regency,
				licensing,
				licensing_number,
				created_at,
				updated_at
			FROM catalog.products WHERE id = $1;`

	row := r.db.QueryRow(ctx, query, id)

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

	return &product, err

}

func (r *ProductRepository) AddProduct(ctx context.Context, product *models.Product) error {
	query := `INSERT INTO catalog.products(
				id,
				category_id,
				name,
				slug,
				description, 
				price,
				stock_quantity,
				weight_grams,
				image_url,
				is_active,
				average_rating,
				benefit,
				composition,
				directions,
				storage_instructions,
				manufacturer,
				marketing_location,
				production_location,
				regency,
				licensing,
				licensing_number,
				created_at,
				updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
    $11, $12, $13, $14, $15, $16, $17, $18,
    $19, $20, $21, NOW(), NOW())`

	ID := uuid.New()

	_, err := r.db.Exec(ctx, query,
		ID,
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
		&product.LicensingNumber)

	return err

}

func (r *ProductRepository) UpdateProductByID(ctx context.Context, id string, product models.Product) error {
	query := `
			UPDATE catalog.products
			SET
				category_id = $2,
				name = $3,
				slug = $4,
				description = $5,
				price = $6,
				stock_quantity = $7,
				weight_grams = $8,
				image_url = $9,
				is_active = $10,
				average_rating = $11,
				benefit = $12,
				composition = $13,
				directions = $14,
				storage_instructions = $15,
				manufacturer = $16,
				marketing_location = $17,
				production_location = $18,
				regency = $19,
				licensing = $20,
				licensing_number = $21,
				updated_at = NOW()
			WHERE id = $1
			`

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

	return err

}

func (r *ProductRepository) DeleteProductByID(ctx context.Context, id string) error {
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
