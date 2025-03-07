package service

import (
	"github.com/bebefabian/orderpack/internal/models"
	"github.com/bebefabian/orderpack/internal/repository"
	"math"
	"sort"
)

// PackServiceImpl implements the PackService interface
type PackServiceImpl struct {
	repo repository.PackRepository
}

type tempSum struct {
	count int // Number of packs used to achieve this sum.
	prev  int // Previous sum.
	pack  int // The pack size used to go from prev to current sum.
}

// CalculatePacks determines the optimal pack distribution
func (s *PackServiceImpl) CalculatePacks(orderQuantity int) (models.CalculateResponse, error) {
	packSizes, err := s.repo.GetPacks()
	if err != nil {
		return models.CalculateResponse{OrderQuantity: orderQuantity, Packs: []models.PackResult{}}, err
	}

	// Sort packs in descending order (largest to smallest)
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	if len(packSizes) == 0 {
		return models.CalculateResponse{OrderQuantity: orderQuantity, Packs: []models.PackResult{}}, err
	}
	maxPack := packSizes[0]
	for _, p := range packSizes {
		if p > maxPack {
			maxPack = p
		}
	}

	//use a tempSum array of size = orderQuantity + maxPack
	limit := orderQuantity + maxPack
	ts := make([]tempSum, limit+1)
	// Initialize tempSum[0]
	ts[0] = tempSum{count: 0, prev: -1, pack: -1}
	// For all other sums, set count to "infinity"
	for s := 1; s <= limit; s++ {
		ts[s] = tempSum{count: math.MaxInt32, prev: -1, pack: -1}
	}

	// Sort packSizes in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	// Fill tempsum table, a hashmap might work ?
	for s := 0; s <= limit; s++ {
		if ts[s].count == math.MaxInt32 {
			continue
		}
		// Try adding each pack size
		for _, pack := range packSizes {
			next := s + pack
			if next > limit {
				continue
			}
			newCount := ts[s].count + 1
			// We want to minimize total shipped first
			// so if we reach the same sum with fewer packs, update.
			if newCount < ts[next].count {
				ts[next] = tempSum{count: newCount, prev: s, pack: pack}
			}
		}
	}

	// Find the minimal total shipped S >= orderQuantity that is reachable.
	bestS := -1
	for s := orderQuantity; s <= limit; s++ {
		if ts[s].count < math.MaxInt32 {
			bestS = s
			break
		}
	}

	// Reconstruct the pack usage from tempSum table.
	usage := make(map[int]int)
	tempS := bestS
	for tempS > 0 {
		cell := ts[tempS]
		usage[cell.pack]++
		tempS = cell.prev
	}

	// Convert the usage map to a slice of PackResult.
	results := []models.PackResult{}
	for _, pack := range packSizes {
		if q, ok := usage[pack]; ok {
			results = append(results, models.PackResult{PackSize: pack, Quantity: q})
		}
	}

	return models.CalculateResponse{
		OrderQuantity: orderQuantity,
		Packs:         results,
	}, nil
}

// NewPackService initializes a new service with the given repository
func NewPackService(repo repository.PackRepository) PackService {
	return &PackServiceImpl{repo: repo}
}

// GetPackSizes returns the current available pack sizes
func (s *PackServiceImpl) GetPackSizes() ([]int, error) {
	results, err := s.repo.GetPacks()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// UpdatePackSizes modifies the available pack sizes dynamically
func (s *PackServiceImpl) UpdatePackSizes(newSizes []int) error {
	err := s.repo.UpdatePacks(newSizes)
	if err != nil {
		return err
	}
	return nil
}
