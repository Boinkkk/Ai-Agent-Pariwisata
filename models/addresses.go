package models

import "github.com/google/uuid"

type Addresses struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Label         string    `json:"label"`
	RecipientName string    `json:"recipient_name"`
	PhoneNumber   string    `json:"phone_number"`
	AddressLine   string    `json:"address_line"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	PostalCode    string    `json:"postal_code"`
	IsMain        bool      `json:"is_main"`
}
