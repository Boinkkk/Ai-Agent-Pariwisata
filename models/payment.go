package models

type Payment struct {
	ID                    string  `json:"id"`
	ExternalTransactionID string  `json:"external_transaction_id,omitempty"`
	PaymentMethod         string  `json:"payment_method"`
	Amount                float64 `json:"amount"`
	PaymentUrl            string  `json:"payment_url,omitempty"`
	PaidAt                string  `json:"paid_at,omitempty"`
}
