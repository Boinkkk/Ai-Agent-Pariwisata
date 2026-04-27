package models

type Courier struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	ServiceType string `json:"service_type"`
}
