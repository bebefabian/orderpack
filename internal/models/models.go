package models

// PackResult represents a single pack in the response
type PackResult struct {
	PackSize int `json:"packSize"`
	Quantity int `json:"quantity"`
}

// CalculateResponse represents the full calculation result
type CalculateResponse struct {
	OrderQuantity int          `json:"orderQuantity"`
	Packs         []PackResult `json:"packs"`
}
