package models

type Order struct {
	ID              string  `json:"id"`
	UserID          string  `json:"user_id"`
	PaymentID       string  `json:"payment_id,omitempty"`
	AddressesID     string  `json:"addresses_id,omitempty"`
	CourierID       string  `json:"courier_id,omitempty"`
	TotalItemsPrice float64 `json:"total_items_price"`
	ShippingCost    float64 `json:"shipping_cost"`
	TotalAmount     float64 `json:"total_amount"`
	Status          string  `json:"status"`
	TrackingNumber  string  `json:"tracking_number,omitempty"`
	CreatedAt       string  `json:"created_at"`
}
