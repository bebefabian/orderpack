package service

import "github.com/bebefabian/orderpack/internal/models"

// PackService defines the interface for managing pack sizes
type PackService interface {
	GetPackSizes() ([]int, error)
	UpdatePackSizes(newSizes []int) error
	CalculatePacks(orderQuantity int) (models.CalculateResponse, error)
}
