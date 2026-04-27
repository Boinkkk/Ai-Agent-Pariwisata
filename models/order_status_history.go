package models

type OrderStatusHistory struct {
	ID        string `json:"id"`
	OrderID   string `json:"order_id"`
	Status    string `json:"status"`
	Notes     string `json:"notes,omitempty"`
	CreatedAt string `json:"created_at"`
}
