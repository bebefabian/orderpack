package service

import "github.com/bebefabian/orderpack/internal/models"

// PackService defines the interface for managing pack sizes
type PackService interface {
	GetPackSizes() []int
	UpdatePackSizes(newSizes []int)
	CalculatePacks(orderQuantity int) models.CalculateResponse
}
