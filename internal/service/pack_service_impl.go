package service

import (
	"github.com/bebefabian/orderpack/internal/models"
	"github.com/bebefabian/orderpack/internal/repository"
	"sort"
)

// PackServiceImpl implements the PackService interface
type PackServiceImpl struct {
	repo repository.PackRepository
}

func (s *PackServiceImpl) CalculatePacks(orderQuantity int) models.CalculateResponse {
	packSizes := s.repo.GetPacks()

	// Sort packs in descending order to prioritize larger packs
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	result := []models.PackResult{}
	remaining := orderQuantity

	for _, pack := range packSizes {
		if remaining == 0 {
			break
		}

		count := remaining / pack
		if count > 0 {
			result = append(result, models.PackResult{PackSize: pack, Quantity: count})
			remaining -= count * pack
		}
	}

	// If there's still a remainder, always use ONE smallest pack
	if remaining > 0 && len(packSizes) > 0 {
		result = append(result, models.PackResult{PackSize: packSizes[len(packSizes)-1], Quantity: 1})
	}

	return models.CalculateResponse{
		OrderQuantity: orderQuantity,
		Packs:         result,
	}
}

// NewPackService initializes a new service with the given repository
func NewPackService(repo repository.PackRepository) PackService {
	return &PackServiceImpl{repo: repo}
}

// GetPackSizes returns the current available pack sizes
func (s *PackServiceImpl) GetPackSizes() []int {
	return s.repo.GetPacks()
}

// UpdatePackSizes modifies the available pack sizes dynamically
func (s *PackServiceImpl) UpdatePackSizes(newSizes []int) {
	s.repo.UpdatePacks(newSizes)
}
