package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID                 uuid.UUID `json:"id"`
	CategoryID         int       `json:"category_id"`
	Name               string    `json:"name"`
	Slug               string    `json:"slug"`
	Description        string    `json:"description"`
	Price              int       `json:"price"`
	StockQuantity      int       `json:"stock_quantity"`
	WeightGrams        int       `json:"weight_grams"`
	ImageURL           string    `json:"image_url"`
	IsActive           bool      `json:"is_active"`
	AverageRating      float32   `json:"average_rating"`
	Benefit            string    `json:"benefit"`
	Composition        string    `json:"composition"`
	Directions         string    `json:"directions"`
	StorageInstruction string    `json:"storage_instructions"`
	Manufacturer       string    `json:"manufacturer"`
	MarketingLocation  string    `json:"marketing_location"`
	ProductionLocation string    `json:"production_location"`
	Regency            string    `json:"regency"`
	Licensing          string    `json:"licensing"`
	LicensingNumber    string    `json:"licensing_number"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
